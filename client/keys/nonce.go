package keys

import (
	"crypto/rand"
	"io"
)

func GenerateNonce() (nonce [24]byte) {
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	return
}
