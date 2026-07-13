package blockchain

import (
	"crypto/sha256"
	"strconv"
)

func (p Payload) Hash(nonce int64) string {
	hash := sha256.Sum256([]byte(strconv.FormatInt(nonce, 10) + p.String()))

	return string(hash[:])
}
