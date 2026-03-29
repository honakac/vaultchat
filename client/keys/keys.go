package keys

import (
	"encoding/gob"
	"errors"
	"os"

	"github.com/honakac/vaultchat/common"
)

func WriteKeys(filepath string, key *common.Keys) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(*key); err != nil {
		panic(err)
	}

	if err := file.Sync(); err != nil {
		panic(err)
	}
}

func ReadKeys(filepath string) (key *common.Keys) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0600)
	if errors.Is(err, os.ErrNotExist) {
		panic("File with keys is not exists!")
	} else if err != nil {
		panic(err)
	}
	defer file.Close()

	key = new(common.Keys)
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(key); err != nil {
		panic(err)
	}

	return
}
