package intergalactictransmission

import (
	"errors"
)

func Transmit(message []byte) []byte {
	bitStream := make([]byte, 0, len(message)*8)
	for _, b := range message {
		for i := 7; i >= 0; i-- {
			bitStream = append(bitStream, (b>>i)&1)
		}
	}
	if pad := (7 - len(bitStream)%7) % 7; pad > 0 {
		bitStream = append(bitStream, make([]byte, pad)...)
	}

	encoded := make([]byte, 0, len(bitStream)/7)
	for i := 0; i < len(bitStream); i += 7 {
		var parity, data byte
		for _, bit := range bitStream[i : i+7] {
			parity ^= bit
			data = (data << 1) | bit
		}
		encoded = append(encoded, (data<<1)|parity)
	}
	return encoded
}

func Decode(message []byte) ([]byte, error) {
	dataBits := make([]byte, 0, len(message)*7)
	for _, b := range message {
		var parity byte
		for i := 0; i < 8; i++ {
			parity ^= (b >> i) & 1
		}
		if parity != 0 {
			return nil, errors.New("wrong parity")
		}
		for i := 7; i >= 1; i-- {
			dataBits = append(dataBits, (b>>i)&1)
		}
	}

	decoded := make([]byte, 0, len(dataBits)/8)
	for i := 0; i+8 <= len(dataBits); i += 8 {
		var by byte
		for _, bit := range dataBits[i : i+8] {
			by = (by << 1) | bit
		}
		decoded = append(decoded, by)
	}
	return decoded, nil
}
