package commands

import (
	"fmt"

	"github.com/honakac/vaultchat/client/keys"
)

func Generate() (key *keys.Keys) {
	fmt.Print("The file already exists, are you sure you want to overwrite it? (y/N): ")
	var agree rune
	fmt.Scanf("%c\n", &agree)

	if agree == 'y' || agree == 'Y' {
		key = keys.GenerateKeys()
		keys.WriteKeys(key)

		fmt.Println("Successfully!")
	} else {
		fmt.Println("Ignored.")
	}

	return
}
