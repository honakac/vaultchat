package main

import (
	"crypto/rand"
	"encoding/gob"
	"errors"
	"os"

	"github.com/honakac/vaultchat/common"
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/crypto/nacl/sign"
)

const (
	FILEKEYS string = "user.keys"
)

func GenerateKeys() (key *common.Keys) {
	publicBox, privateBox, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	publicSign, privateSign, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	return &common.Keys{
		PublicBox:   publicBox,
		PrivateBox:  privateBox,
		PublicSign:  publicSign,
		PrivateSign: privateSign,
	}
}

func WriteKeys(key *common.Keys) {
	file, err := os.OpenFile(FILEKEYS, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(*key); err != nil {
		panic(err)
	}

	if err := file.Sync(); err != nil {
		panic(err)
	}
}

func ReadKeys() (key *common.Keys) {
	file, err := os.OpenFile(FILEKEYS, os.O_RDONLY, 0600)
	if errors.Is(err, os.ErrNotExist) {
		panic("File with keys is not exists!")
	} else if err != nil {
		panic(err)
	}
	defer file.Close()

	key = new(common.Keys)
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(key); err != nil {
		panic(err)
	}

	return
}
