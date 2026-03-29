package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v3"
	"github.com/honakac/vaultchat/client/keys"
	"github.com/honakac/vaultchat/common"
	"github.com/honakac/vaultchat/relay/database"
)

const (
	DEFAULT_ADDR string = ":4269"
	API_VERSION  string = "/v1"
)

var Key *common.Keys
var KeyId string
var Db database.Database

func main() {
	if _, err := os.Stat("relay.keys"); errors.Is(err, os.ErrNotExist) {
		Key = common.GenerateKeys()
		keys.WriteKeys("relay.keys", Key)

		fmt.Println("Keys is successfully generated!")
		fmt.Println()
	} else {
		Key = keys.ReadKeys("relay.keys")
	}

	KeyId = common.PackID(Key.PublicBox, Key.PublicSign)
	fmt.Println("Your relay id:", KeyId)

	addr := DEFAULT_ADDR
	if len(os.Args) == 2 {
		addr = os.Args[1]
	}

	Db.Init("relay.db")

	app := fiber.New(fiber.Config{})

	app.Get(API_VERSION+"/info", func(c fiber.Ctx) error {
		return c.SendString(KeyId)
	})

	app.Post(API_VERSION+"/send_message/:sender_id", func(c fiber.Ctx) error {
		body := c.Body()
		senderId := c.Params("sender_id")

		decrypted, err := common.DecryptById(Key, senderId, body)
		if err != nil {
			log.Errorf("Failed to decrypt message: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to decrypt message",
			})
		}

		var req database.SendMessageRequest
		if err := json.Unmarshal(decrypted, &req); err != nil {
			log.Errorf("Failed to unmarshal decrypted message: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to unmarshal decrypted message",
			})
		}

		return Db.AddInboxMessage(c, req)
	})

	app.Get(API_VERSION+"/fetch_messages/:receiver_id/:last_cuid", func(c fiber.Ctx) error {
		receiverId := c.Params("receiver_id")
		lastCuid := c.Params("last_cuid")

		return Db.GetInboxMessages(Key, c, receiverId, lastCuid)
	})

	log.Fatal(app.Listen(addr, fiber.ListenConfig{
		DisableStartupMessage: false,
	}))
}
