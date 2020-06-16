package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func ApplySha256(record string) string {
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
