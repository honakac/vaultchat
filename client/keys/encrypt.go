package keys

import (
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/crypto/nacl/sign"
)

func EncryptByPublicKey(key *Keys, recipientPublicKey *[32]byte, message string) (encrypted []byte) {
	nonce := GenerateNonce()
	signed := sign.Sign(nil, []byte(message), key.PrivateSign)
	encrypted = box.Seal(nonce[:], signed, &nonce, recipientPublicKey, key.PrivateBox)

	return
}

func EncryptById(key *Keys, recipientID string, message string) ([]byte, error) {
	recipientBox, _, err := ExtractID(recipientID)
	if err != nil {
		return nil, err
	}

	return EncryptByPublicKey(key, (*[32]byte)(recipientBox), message), nil
}
