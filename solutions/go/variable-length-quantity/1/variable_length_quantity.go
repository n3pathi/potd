package variablelengthquantity

import "errors"

func EncodeVarint(input []uint32) []byte {
	if len(input) == 0 {
		return []byte{}
	}

	var result []byte
	for _, v := range input {
		result = append(result, VLEncoder(v)...)
	}
	return result
}

func VLEncoder(v uint32) []byte {
	// Peel off 7-bit groups starting from the low end; the first group (least
	// significant) never gets the continuation bit since it's written last.
	buffer := []byte{byte(v & 0x7F)}
	v >>= 7
	for v > 0 {
		// Every group after the first is a *more* significant chunk, so it
		// needs the continuation bit (0x80) set to mark "more bytes follow".
		buffer = append(buffer, byte(v&0x7F)|0x80)
		v >>= 7
	}

	// buffer is currently least-significant-group first; VLQ transmits
	// most-significant-group first, so reverse it in place before returning.
	for i, j := 0, len(buffer)-1; i < j; i, j = i+1, j-1 {
		buffer[i], buffer[j] = buffer[j], buffer[i]
	}
	return buffer
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var output []uint32
	var num uint32
	for i, b := range input {
		// Groups arrive most-significant first, so each new 7-bit chunk
		// slots in at the bottom and pushes the accumulated bits left.
		num = num<<7 | uint32(b&0x7F)
		if b&0x80 == 0 {
			// Continuation bit unset: this byte completes the current value.
			output = append(output, num)
			num = 0 // reset the accumulator for the next value in the stream
		} else if i == len(input)-1 {
			// Continuation bit set on the last byte means the stream was
			// truncated mid-value; there's no next byte to complete it.
			return nil, errors.New("incomplete sequence")
		}
	}
	return output, nil
}
