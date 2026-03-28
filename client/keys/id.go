package keys

import (
	"github.com/btcsuite/btcutil/base58"
)

func ExtractID(id string) (*[32]byte, *[32]byte, error) {
	boxKey := new([32]byte)
	signKey := new([32]byte)

	key := base58.Decode(id)

	copy(boxKey[:], key[:32])
	copy(signKey[:], key[32:])

	return boxKey, signKey, nil
}

func PackID(boxKey *[32]byte, signKey *[32]byte) (id string) {
	id = base58.Encode(append(boxKey[:], signKey[:]...))

	return
}
