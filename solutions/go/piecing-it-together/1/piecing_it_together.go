package piecingittogether

import (
	"errors"
	"math"
)

type PuzzleDetails struct {
	Pieces      int
	Border      int
	Inside      int
	Rows        int
	Columns     int
	AspectRatio float64
	Format      string
}

const epsilon = 1e-9

// puzzle tracks partial jigsaw data as it is progressively solved. A field
// counts as known only once its has* flag is set, so a legitimately zero
// value (e.g. Inside == 0 for an all-border puzzle) is never mistaken for
// "not yet provided".
type puzzle struct {
	pieces, border, inside, rows, columns int
	aspectRatio                           float64
	format                                string

	hasPieces, hasBorder, hasInside, hasRows, hasColumns bool
	hasAspectRatio, hasFormat                            bool
}

func newPuzzle(d PuzzleDetails) *puzzle {
	return &puzzle{
		pieces: d.Pieces, border: d.Border, inside: d.Inside,
		rows: d.Rows, columns: d.Columns,
		aspectRatio: d.AspectRatio, format: d.Format,

		hasPieces:      d.Pieces > 0,
		hasBorder:      d.Border > 0,
		hasInside:      d.Inside > 0,
		hasRows:        d.Rows > 0,
		hasColumns:     d.Columns > 0,
		hasAspectRatio: d.AspectRatio > 0,
		hasFormat:      d.Format != "",
	}
}

func setInt(cur *int, has *bool, val int) error {
	if *has {
		if *cur != val {
			return errors.New("Contradictory data")
		}
		return nil
	}
	*cur, *has = val, true
	return nil
}

func setFloat(cur *float64, has *bool, val float64) error {
	if *has {
		if math.Abs(*cur-val) > epsilon {
			return errors.New("Contradictory data")
		}
		return nil
	}
	*cur, *has = val, true
	return nil
}

func (p *puzzle) setFormat(val string) error {
	if p.hasFormat {
		if p.format != val {
			return errors.New("Contradictory data")
		}
		return nil
	}
	p.format, p.hasFormat = val, true
	return nil
}

func round(f float64) int {
	return int(math.Round(f))
}

// Each step below fills in fields implied by ones already known, using the
// jigsaw relationships between Rows/Columns (R, C) and the rest:
// Pieces = R*C, Border = 2*(R+C)-4, Inside = (R-2)*(C-2), AspectRatio = C/R.
// solve() runs them to a fixed point, so it doesn't matter which fields the
// caller happened to supply.

func (p *puzzle) fillPiecesFromBorderInside() error {
	if !p.hasBorder || !p.hasInside {
		return nil
	}
	return setInt(&p.pieces, &p.hasPieces, p.border+p.inside)
}

func (p *puzzle) fillBorderFromPiecesInside() error {
	if !p.hasPieces || !p.hasInside {
		return nil
	}
	return setInt(&p.border, &p.hasBorder, p.pieces-p.inside)
}

func (p *puzzle) fillInsideFromPiecesBorder() error {
	if !p.hasPieces || !p.hasBorder {
		return nil
	}
	return setInt(&p.inside, &p.hasInside, p.pieces-p.border)
}

func (p *puzzle) fillFromRowsColumns() error {
	if !p.hasRows || !p.hasColumns {
		return nil
	}
	if err := setInt(&p.pieces, &p.hasPieces, p.rows*p.columns); err != nil {
		return err
	}
	if err := setInt(&p.border, &p.hasBorder, 2*(p.rows+p.columns)-4); err != nil {
		return err
	}
	if err := setInt(&p.inside, &p.hasInside, (p.rows-2)*(p.columns-2)); err != nil {
		return err
	}
	return setFloat(&p.aspectRatio, &p.hasAspectRatio, float64(p.columns)/float64(p.rows))
}

func (p *puzzle) fillColumnsFromRowsAspectRatio() error {
	if !p.hasRows || !p.hasAspectRatio || p.hasColumns {
		return nil
	}
	return setInt(&p.columns, &p.hasColumns, round(float64(p.rows)*p.aspectRatio))
}

func (p *puzzle) fillRowsFromColumnsAspectRatio() error {
	if !p.hasColumns || !p.hasAspectRatio || p.hasRows {
		return nil
	}
	return setInt(&p.rows, &p.hasRows, round(float64(p.columns)/p.aspectRatio))
}

func (p *puzzle) fillRowsColumnsFromPiecesAspectRatio() error {
	if !p.hasPieces || !p.hasAspectRatio || (p.hasRows && p.hasColumns) {
		return nil
	}
	r := round(math.Sqrt(float64(p.pieces) / p.aspectRatio))
	c := round(math.Sqrt(float64(p.pieces) * p.aspectRatio))
	if err := setInt(&p.rows, &p.hasRows, r); err != nil {
		return err
	}
	return setInt(&p.columns, &p.hasColumns, c)
}

// fillRowsColumnsFromSquareInside covers the square case, where
// Inside = (side-2)^2 lets side be recovered from Inside alone.
func (p *puzzle) fillRowsColumnsFromSquareInside() error {
	if !p.hasAspectRatio || p.aspectRatio != 1 || !p.hasInside || p.hasRows {
		return nil
	}
	side := round(math.Sqrt(float64(p.inside))) + 2
	if err := setInt(&p.rows, &p.hasRows, side); err != nil {
		return err
	}
	return setInt(&p.columns, &p.hasColumns, side)
}

// fillRowsColumnsFromPiecesBorder covers the general case: Pieces = R*C and
// Border = 2*(R+C)-4 pin {R, C} down as the roots of
// x^2 - (Border/2+2)*x + Pieces = 0; Format picks which root is which.
func (p *puzzle) fillRowsColumnsFromPiecesBorder() error {
	if !p.hasPieces || !p.hasBorder || (p.hasRows && p.hasColumns) {
		return nil
	}
	sum := p.border/2 + 2
	disc := sum*sum - 4*p.pieces
	if disc < 0 {
		return nil
	}
	sq := round(math.Sqrt(float64(disc)))
	x1, x2 := (sum+sq)/2, (sum-sq)/2
	big, small := max(x1, x2), min(x1, x2)

	rows, columns := 0, 0
	switch {
	case x1 == x2:
		rows, columns = x1, x1
	case p.hasFormat && p.format == "portrait":
		rows, columns = big, small
	case p.hasFormat && p.format == "landscape":
		rows, columns = small, big
	default:
		return nil // ambiguous orientation without a Format hint
	}
	if err := setInt(&p.rows, &p.hasRows, rows); err != nil {
		return err
	}
	return setInt(&p.columns, &p.hasColumns, columns)
}

func (p *puzzle) fillAspectRatioFromFormat() error {
	if p.format != "square" {
		return nil
	}
	return setFloat(&p.aspectRatio, &p.hasAspectRatio, 1.0)
}

func (p *puzzle) fillFormatFromAspectRatio() error {
	if !p.hasAspectRatio || p.hasFormat {
		return nil
	}
	switch {
	case p.aspectRatio > 1:
		return p.setFormat("landscape")
	case p.aspectRatio == 1:
		return p.setFormat("square")
	default:
		return p.setFormat("portrait")
	}
}

// solve runs every fill step to a fixed point, or until one reports a
// contradiction. Each pass can only turn unknown fields into known ones (or
// error out), so the loop is bounded by the number of fields there are to
// learn.
func (p *puzzle) solve() error {
	steps := []func() error{
		p.fillPiecesFromBorderInside,
		p.fillBorderFromPiecesInside,
		p.fillFromRowsColumns,
		p.fillColumnsFromRowsAspectRatio,
		p.fillInsideFromPiecesBorder,
		p.fillRowsFromColumnsAspectRatio,
		p.fillRowsColumnsFromPiecesAspectRatio,
		p.fillRowsColumnsFromSquareInside,
		p.fillRowsColumnsFromPiecesBorder,
		p.fillAspectRatioFromFormat,
		p.fillFormatFromAspectRatio,
	}
	for range len(steps) {
		before := *p
		for _, step := range steps {
			if err := step(); err != nil {
				return err
			}
		}
		if *p == before {
			break
		}
	}
	return nil
}

func JigsawData(details PuzzleDetails) (PuzzleDetails, error) {
	p := newPuzzle(details)
	if err := p.solve(); err != nil {
		return PuzzleDetails{}, err
	}
	if !p.hasPieces || !p.hasBorder || !p.hasInside || !p.hasRows ||
		!p.hasColumns || !p.hasAspectRatio || !p.hasFormat {
		return PuzzleDetails{}, errors.New("Insufficient data")
	}
	return PuzzleDetails{
		Pieces:      p.pieces,
		Border:      p.border,
		Inside:      p.inside,
		Rows:        p.rows,
		Columns:     p.columns,
		AspectRatio: p.aspectRatio,
		Format:      p.format,
	}, nil
}
