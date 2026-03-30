package common

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
)

const MinPacketSize = 2048
const AlignSize = 1024

func Padding(payload []byte) ([]byte, error) {
	dataLen := len(payload)
	headerSize := 8
	currentSize := dataLen + headerSize

	targetSize := MinPacketSize
	if currentSize > MinPacketSize {
		targetSize = ((currentSize + AlignSize - 1) / AlignSize) * AlignSize
	}

	buffer := make([]byte, targetSize)

	binary.BigEndian.PutUint32(buffer[0:4], uint32(dataLen))
	binary.BigEndian.PutUint32(buffer[4:8], 0)

	copy(buffer[8:], payload)

	paddingStart := headerSize + dataLen
	if paddingStart < targetSize {
		_, err := rand.Read(buffer[paddingStart:])
		if err != nil {
			return nil, err
		}
	}

	return buffer, nil
}

func Unpadding(buffer []byte) ([]byte, error) {
	if len(buffer) < 8 {
		return nil, errors.New("packet too small")
	}

	dataLen := binary.BigEndian.Uint32(buffer[0:4])

	if int(dataLen)+8 > len(buffer) {
		return nil, errors.New("corrupted header: length mismatch")
	}

	return buffer[8 : 8+dataLen], nil
}
