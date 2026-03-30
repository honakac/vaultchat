package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/honakac/vaultchat/client/api"
	"github.com/honakac/vaultchat/common"
	"github.com/honakac/vaultchat/relay/database"
	"github.com/nrednav/cuid2"
)

func Send(key *common.Keys, keyId string, url string, id string, message string) {
	fmt.Println("Getting relay keys...")

	relayId, err := api.GetInfo(url)
	if err != nil {
		fmt.Printf("Failed to get relay info: %s\n", err)
		return
	}

	fmt.Printf("Successfully, Relay id: %s\n", relayId)

	fmt.Println("Encrypting message to recipient...")

	encryptedMessage, err := common.EncryptById(key, id, message)
	if err != nil {
		fmt.Printf("Failed to encrypt message: %s\n", err)
		return
	}

	fmt.Println("Encrypting message to relay...")

	relayMessage, err := json.Marshal(database.SendMessageRequest{
		Cuid:         cuid2.Generate(),
		ReceiverAddr: id,
		SenderAddr:   keyId,
		Payload:      encryptedMessage,
	})
	if err != nil {
		fmt.Printf("Failed to marshal message: %s\n", err)
		return
	}

	encryptedRelay, err := common.EncryptById(key, relayId, string(relayMessage))
	if err != nil {
		fmt.Printf("Failed to encrypt message to relay: %s\n", err)
		return
	}

	fmt.Println("Sending message to relay...")

	_, err = http.Post(url+"/v1/send_message/"+keyId, "application/octet-stream", bytes.NewReader(encryptedRelay))
	if err != nil {
		fmt.Printf("Failed to send message: %s\n", err)
		return
	}

	fmt.Println("Message sent successfully!")
}
