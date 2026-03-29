package common

import (
	"errors"

	"golang.org/x/crypto/nacl/box"
	"golang.org/x/crypto/nacl/sign"
)

func DecryptByPublicKey(key *Keys, packet []byte, senderPublicBox *[32]byte, senderPublicSign *[32]byte) ([]byte, error) {
	if len(packet) < 24 {
		return nil, errors.New("packet too short")
	}

	var nonce [24]byte
	copy(nonce[:], packet[:24])

	decrypted, ok := box.Open(nil, packet[24:], &nonce, senderPublicBox, key.PrivateBox)
	if !ok {
		return nil, errors.New("failed to decrypt (wrong keys or corrupted data)")
	}

	msg, ok := sign.Open(nil, decrypted, senderPublicSign)
	if !ok {
		return nil, errors.New("invalid signature (message is not from the claimed sender)")
	}

	return msg, nil
}

func DecryptById(key *Keys, senderId string, packet []byte) ([]byte, error) {
	senderBox, senderSign, err := ExtractID(senderId)
	if err != nil {
		return nil, err
	}

	return DecryptByPublicKey(key, packet, (*[32]byte)(senderBox), (*[32]byte)(senderSign))
}
