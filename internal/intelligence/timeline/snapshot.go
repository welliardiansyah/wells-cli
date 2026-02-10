package timeline

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"
)

type Snapshot struct {
	Entity    string    `json:"entity"`
	Hash      string    `json:"hash"`
	Timestamp time.Time `json:"timestamp"`
}

func CreateSnapshot(entity, file string) (*Snapshot, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(data)
	return &Snapshot{
		Entity:    entity,
		Hash:      fmt.Sprintf("%x", hash),
		Timestamp: time.Now(),
	}, nil
}
