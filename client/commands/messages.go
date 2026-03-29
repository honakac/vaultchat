package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/honakac/vaultchat/client/utils"
	"github.com/honakac/vaultchat/common"
	"github.com/honakac/vaultchat/relay/database"
)

type GetMessageResponse struct {
	messages []database.InboxMessage
}

func Messages(key *common.Keys, keyId string, url string) {
	relayId, err := utils.GetInfo(url)
	if err != nil {
		fmt.Printf("Failed to get relay info: %s\n", err)
		return
	}

	relayResponse, err := http.Get(url + "/v1/fetch_messages/" + keyId + "/0")
	if err != nil {
		fmt.Println("Failed to get relay inbox messages:", err)
		return
	}

	relayRawMessages, err := io.ReadAll(relayResponse.Body)
	if err != nil {
		fmt.Printf("Failed to read body data: %s\n", err)
		return
	}

	relayRawDecrypted, err := common.DecryptById(key, relayId, relayRawMessages)
	if err != nil {
		fmt.Printf("Failed to decrypt relay response: %s\n", err)
		return
	}

	var relayMessages database.GetMessagesResponse
	if err := json.Unmarshal(relayRawDecrypted, &relayMessages); err != nil {
		fmt.Printf("Failed to unmarshal JSON: %s\n", err)
		return
	}

	for _, msg := range relayMessages.Messages {
		decrypted, err := common.DecryptById(key, msg.SenderAddr, msg.Payload)
		if err != nil {
			fmt.Printf("Failed to decrypt message: %s\n", err)
			continue
		}

		fmt.Printf("Message from %s: %s\n", msg.SenderAddr, string(decrypted))
	}
}
