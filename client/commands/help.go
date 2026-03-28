package commands

import "fmt"

func Help() {
	fmt.Println("generate                              Generate new keys")
	fmt.Println("encrypt recipientID message           Encrypt message")
	fmt.Println("decrypt senderID message              Decrypt message")
	fmt.Println("exit                                  Just exit")
}
