package commands

import (
	"fmt"

	"github.com/honakac/vaultchat/client/api"
	"github.com/honakac/vaultchat/common"
)

func Messages(key *common.Keys, keyId string, url string) {
	messages, err := api.Messages(key, keyId, url, "0")
	if err != nil {
		fmt.Printf("Failed to get messages: %s\n", err)
		return
	}

	for _, msg := range messages {
		fmt.Printf("Message from %s: %s\n", msg.SenderAddr, msg.Message)
	}
}
