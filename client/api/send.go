package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/honakac/vaultchat/common"
	"github.com/honakac/vaultchat/relay/database"
	"github.com/nrednav/cuid2"
)

func Send(key *common.Keys, keyId string, url string, id string, message string) error {
	relayId, err := GetInfo(url)
	if err != nil {
		return err
	}

	encryptedMessage, err := common.EncryptById(key, id, message)
	if err != nil {
		return err
	}

	fmt.Println("Encrypting message to relay...")

	relayMessage, err := json.Marshal(database.SendMessageRequest{
		Cuid:         cuid2.Generate(),
		ReceiverAddr: id,
		SenderAddr:   keyId,
		Payload:      encryptedMessage,
	})
	if err != nil {
		return err
	}

	encryptedRelay, err := common.EncryptById(key, relayId, string(relayMessage))
	if err != nil {
		return err
	}

	_, err = http.Post(url+"/v1/send_message/"+keyId, "application/octet-stream", bytes.NewReader(encryptedRelay))
	if err != nil {
		return err
	}

	return nil
}
