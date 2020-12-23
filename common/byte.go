package common

import (
	"encoding/binary"
)

func UintToVarLenBytes(l uint64) []byte {
	b := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(b, l)
	trimTrailingZeroByte(&b)
	return b
}

func trimTrailingZeroByte(b *[]byte) {
	for len(*b) >= 0 {
		if (*b)[len(*b)-1] != 0 {
			break
		}
		*b = (*b)[:len(*b)-1]
	}
}
