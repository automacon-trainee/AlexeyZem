package internal

import (
	"encoding/binary"
	"hash"
)

func Hash(key string, hasher *hash.Hash) int {
	(*hasher).Reset()
	(*hasher).Write([]byte(key))
	return int(binary.BigEndian.Uint16((*hasher).Sum(nil)))
}
