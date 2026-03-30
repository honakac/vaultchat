package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/honakac/vaultchat/common"
	"github.com/honakac/vaultchat/relay/database"
)

type GetMessageResponse struct {
	messages []database.InboxMessage
}

type InboxDecryptedMessage struct {
	SenderAddr string
	Message    string
}

func Messages(key *common.Keys, keyId string, url string, lastCuid string) ([]InboxDecryptedMessage, error) {
	var inboxDecryptedMessages []InboxDecryptedMessage

	relayId, err := GetInfo(url)
	if err != nil {
		return nil, err
	}

	relayResponse, err := http.Get(url + "/v1/fetch_messages/" + keyId + "/" + lastCuid)
	if err != nil {
		return nil, err
	}

	relayRawMessages, err := io.ReadAll(relayResponse.Body)
	if err != nil {
		return nil, err
	}

	relayRawDecrypted, err := common.DecryptById(key, relayId, relayRawMessages)
	if err != nil {
		return nil, err
	}

	var relayMessages database.GetMessagesResponse
	if err := json.Unmarshal(relayRawDecrypted, &relayMessages); err != nil {
		return nil, err
	}

	for _, msg := range relayMessages.Messages {
		decrypted, err := common.DecryptById(key, msg.SenderAddr, msg.Payload)
		if err != nil {
			continue
		}

		inboxDecryptedMessages = append(inboxDecryptedMessages, InboxDecryptedMessage{
			SenderAddr: msg.SenderAddr,
			Message:    string(decrypted),
		})
	}

	return inboxDecryptedMessages, nil
}
