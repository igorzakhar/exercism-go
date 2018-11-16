package variablelengthquantity

import (
	"errors"
)

func EncodeVarint(input []uint32) (bytes []byte) {
	for _, i := range input {
		bytes = append(bytes, encodeInt(i)...)
	}
	return
}

func encodeInt(value uint32) (bytes []byte) {
	if value>>7 == 0 {
		return []byte{byte(value)}
	}

	for value > 0 {
		bytes = append([]byte{byte((value & 0x7f) | 0x80)}, bytes...)
		value >>= 7
	}
	bytes[len(bytes)-1] = bytes[len(bytes)-1] & 0x7f
	return
}

func DecodeVarint(input []byte) ([]uint32, error) {
	if input[len(input)-1]&0x80 > 0 {
		return nil, errors.New("incomplete sequence")
	}

	if len(input) == 0 {
		return []uint32{0}, nil
	}

	out := make([]uint32, 0)
	var x uint32

	for _, b := range input {
		x <<= 7
		x |= (uint32(b) & 0x7f)
		if (b & 0x80) == 0 {
			out = append(out, x)
			x = 0
		}
	}
	return out, nil
}
