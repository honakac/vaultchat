package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/honakac/vaultchat/client/commands"
	"github.com/honakac/vaultchat/client/keys"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect database")
	// }

	var key *keys.Keys
	var keyId string

	if _, err := os.Stat("user.keys"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("The key file has not been created, we are generating...")
		fmt.Println("In case of creating new keys, there is a command 'generate'")

		key = keys.GenerateKeys()
		keys.WriteKeys(key)

		fmt.Println("Successfully generated!")
		fmt.Println()
	} else {
		key = keys.ReadKeys()
	}

	keyId = keys.PackID(key.PublicBox, key.PublicSign)
	fmt.Println("Your id:", keyId)

	for {
		fmt.Print("> ")

		if scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			words := strings.Split(line, " ")

			switch words[0] {
			case "encrypt":
				if len(words) < 3 {
					fmt.Println("usage: encrypt recipientID message")
					break
				}
				commands.Encrypt(key, words[1], strings.Join(words[2:], " "))

			case "decrypt":
				if len(words) < 3 {
					fmt.Println("usage: decrypt senderId message")
					break
				}
				commands.Decrypt(key, words[1], strings.Join(words[2:], " "))

			case "generate":
				if newKey := commands.Generate(); newKey != nil {
					key = newKey
					keyId = keys.PackID(key.PublicBox, key.PublicSign)
					fmt.Println("Your new id:", keyId)
				}

			case "help":
				commands.Help()
			case "exit", "quit":
				os.Exit(0)
			default:
				fmt.Println("Unknown command, type 'help'")
			}
		} else {
			os.Exit(1)
		}
	}
}
