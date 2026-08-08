package robot

import "fmt"

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.

const (
	N Dir = iota
	E
	S
	W
)

// turnRight and turnLeft are the single source of truth for how a
// direction changes on a turn; Step1 and Step2 both call them instead
// of each re-deriving the N/E/S/W cycle.
func turnRight(d Dir) Dir {
	switch d {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	default: // W
		return N
	}
}

func turnLeft(d Dir) Dir {
	switch d {
	case N:
		return W
	case W:
		return S
	case S:
		return E
	default: // E
		return N
	}
}

// dirDelta is the single source of truth for how far a step in
// direction d moves you along each axis.
func dirDelta(d Dir) (dx, dy int) {
	switch d {
	case N:
		return 0, 1
	case E:
		return 1, 0
	case S:
		return 0, -1
	default: // W
		return -1, 0
	}
}

func Right() {
	Step1Robot.Dir = turnRight(Step1Robot.Dir)
}

func Left() {
	Step1Robot.Dir = turnLeft(Step1Robot.Dir)
}

func Advance() {
	dx, dy := dirDelta(Step1Robot.Dir)
	Step1Robot.X += dx
	Step1Robot.Y += dy
}

func (d Dir) String() string {
	switch d {
	case N:
		return "North"
	case E:
		return "East"
	case S:
		return "South"
	case W:
		return "West"
	default:
		panic("unknown direction")
	}
}

// Step 2
// Define Action type here.
type Action int

const (
	START Action = iota
	R
	L
	A
)

func StartRobot(command chan Command, action chan Action) {
	for cmd := range command {
		switch cmd {
		case 'R':
			action <- R
		case 'L':
			action <- L
		case 'A':
			action <- A
		}
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for a := range action {
		switch a {
		case R:
			robot.Dir = turnRight(robot.Dir)
		case L:
			robot.Dir = turnLeft(robot.Dir)
		case A:
			dx, dy := dirDelta(robot.Dir)
			next := Pos{robot.Pos.Easting + RU(dx), robot.Pos.Northing + RU(dy)}
			within := next.Easting >= extent.Min.Easting && next.Easting <= extent.Max.Easting &&
				next.Northing >= extent.Min.Northing && next.Northing <= extent.Max.Northing
			if within {
				robot.Pos = next
			}
		}
	}
	report <- robot
}

// Step 3
// Define Action3 type here.
type Action3 struct {
	Name string
	Cmd  byte // 'R', 'L', or 'A'
	Done bool // true on the final message from a robot's script
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	for i := range script {
		switch c := script[i]; c {
		case 'R', 'L', 'A':
			action <- Action3{Name: name, Cmd: c}
		default:
			log <- fmt.Sprintf("%s: unknown command %q", name, c)
			action <- Action3{Name: name, Done: true}
			return
		}
	}
	action <- Action3{Name: name, Done: true}
}

func within(extent Rect, p Pos) bool {
	return p.Easting >= extent.Min.Easting && p.Easting <= extent.Max.Easting &&
		p.Northing >= extent.Min.Northing && p.Northing <= extent.Max.Northing
}

func indexByName(robots []Step3Robot, name string) int {
	for i, r := range robots {
		if r.Name == name {
			return i
		}
	}
	return -1
}

func occupiedBy(robots []Step3Robot, self int, p Pos) bool {
	for i, r := range robots {
		if i != self && r.Pos == p {
			return true
		}
	}
	return false
}

// validateRoster logs the placement problems called out in the
// instructions: a missing name, a name reused by another robot, a
// robot placed outside the room, and two robots placed on top of
// each other.
func validateRoster(extent Rect, robots []Step3Robot, log chan string) {
	seen := map[string]bool{}
	for i, r := range robots {
		switch {
		case r.Name == "":
			log <- "a robot was placed with no name"
		case seen[r.Name]:
			log <- fmt.Sprintf("robot name %q is used more than once", r.Name)
		}
		seen[r.Name] = true

		if !within(extent, r.Pos) {
			log <- fmt.Sprintf("robot %q was placed outside the room", r.Name)
		}
		for j := 0; j < i; j++ {
			if robots[j].Pos == r.Pos {
				log <- fmt.Sprintf("robots %q and %q were placed at the same position", robots[j].Name, r.Name)
				break
			}
		}
	}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	validateRoster(extent, robots, log)

	done := 0
	for done < len(robots) {
		a := <-action
		if a.Done {
			done++
			continue
		}

		idx := indexByName(robots, a.Name)
		if idx < 0 {
			log <- fmt.Sprintf("action from unknown robot %q", a.Name)
			continue
		}
		r := &robots[idx]

		switch a.Cmd {
		case 'R':
			r.Dir = turnRight(r.Dir)
		case 'L':
			r.Dir = turnLeft(r.Dir)
		case 'A':
			dx, dy := dirDelta(r.Dir)
			next := Pos{r.Pos.Easting + RU(dx), r.Pos.Northing + RU(dy)}
			switch {
			case !within(extent, next):
				log <- fmt.Sprintf("%s attempted to advance into a wall", a.Name)
			case occupiedBy(robots, idx, next):
				log <- fmt.Sprintf("%s attempted to advance into another robot", a.Name)
			default:
				r.Pos = next
			}
		}
	}
	rep <- robots
}
