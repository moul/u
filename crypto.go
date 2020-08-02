package u

import (
	"crypto/sha1"
	"fmt"
)

func Sha1(data []byte) []byte {
	h := sha1.New()
	_, _ = h.Write(data)
	return h.Sum(nil)
}

func Sha1Hex(data []byte) string {
	return fmt.Sprintf("%x", Sha1(data))
}
