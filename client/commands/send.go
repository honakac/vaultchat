package commands

import (
	"fmt"

	"github.com/honakac/vaultchat/client/api"
	"github.com/honakac/vaultchat/common"
)

func Send(key *common.Keys, keyId string, url string, id string, message string) {
	err := api.Send(key, keyId, url, id, message)
	if err != nil {
		fmt.Printf("Failed to send message: %s\n", err)
	} else {
		fmt.Println("Message sent successfully!")
	}
}
