package Helper

import (
	"crypto/rand"
	"io"
)

// UUIDGenerator creates unique id
type UUIDGenerator struct {
}

// CreateUUID returns unique id
func (uuid *UUIDGenerator) CreateUUID() ([]byte, error) {
	newUUID := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, newUUID)
	if n != len(newUUID) || err != nil {
		return nil, err
	}
	// variant bits
	newUUID[8] = newUUID[8]&^0xc0 | 0x80
	// version 4 (pseudo-random)
	newUUID[6] = newUUID[6]&^0xf0 | 0x40
	return newUUID, nil
}
