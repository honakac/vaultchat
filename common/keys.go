package common

import (
	"crypto/rand"

	"golang.org/x/crypto/nacl/box"
	"golang.org/x/crypto/nacl/sign"
)

type Keys struct {
	PublicBox   *[32]byte
	PrivateBox  *[32]byte
	PublicSign  *[32]byte
	PrivateSign *[64]byte
}

func GenerateKeys() (key *Keys) {
	publicBox, privateBox, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	publicSign, privateSign, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	return &Keys{
		PublicBox:   publicBox,
		PrivateBox:  privateBox,
		PublicSign:  publicSign,
		PrivateSign: privateSign,
	}
}
