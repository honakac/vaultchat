package commands

import (
	"encoding/base64"
	"fmt"
	"github.com/honakac/vaultchat/common"
)

func Encrypt(key *common.Keys, id string, message string) {
	data, err := common.EncryptById(key, id, message)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(base64.StdEncoding.EncodeToString(data))
	}
}
