package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(content []byte) string {
	h := sha256.New()
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}
