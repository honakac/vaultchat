package utils

import (
	"io"
	"net/http"
)

func GetInfo(url string) (string, error) {
	relayInfo, err := http.Get(url + "/v1/info")
	if err != nil {
		return "", err
	}

	relayRawId, err := io.ReadAll(relayInfo.Body)
	if err != nil {
		return "", err
	}

	return string(relayRawId), nil
}
