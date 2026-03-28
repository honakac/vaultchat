package commands

import (
	"encoding/base64"
	"fmt"

	"github.com/honakac/vaultchat/client/keys"
)

func Encrypt(key *keys.Keys, id string, message string) {
	data, err := keys.EncryptById(key, id, message)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(base64.StdEncoding.EncodeToString(data))
	}
}
