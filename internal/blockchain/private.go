package blockchain

import (
	"crypto/sha256"
	"strconv"
)

func (p Payload) Hash(nonce int) string {
	hash := sha256.Sum256([]byte(strconv.Itoa(nonce) + p.String()))

	return string(hash[:])
}
