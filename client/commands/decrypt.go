package commands

import (
	"encoding/base64"
	"fmt"

	"github.com/honakac/vaultchat/common"
)

func Decrypt(key *common.Keys, id string, messageBase64 string) {
	message, err := base64.StdEncoding.DecodeString(messageBase64)
	if err != nil {
		panic(err)
	}

	data, err := common.DecryptById(key, id, message)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}
}
