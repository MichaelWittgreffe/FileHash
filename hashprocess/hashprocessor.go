package hashprocess

import (
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
)

//HashProcessor implements HashProcessor
type HashProcessor struct {
	filepath string
	hasher   hash.Hash
}

//Process performs a hash function on the given filepath and returns the hash
func (p *HashProcessor) Process() (string, error) {
	if len(p.filepath) <= 0 {
		return "", errors.New("Filepath Is Empty")
	}

	file, err := os.Open(p.filepath)
	if err != nil {
		return "", fmt.Errorf("Unable To Open File %s: %s", p.filepath, err.Error())
	}

	_, err = io.Copy(p.hasher, file)
	if err != nil {
		return "", err
	}

	hashBytes := p.hasher.Sum(nil)
	if len(hashBytes) <= 0 {
		return "", fmt.Errorf("hasher.Sum Is Nil")
	}

	return hex.EncodeToString(hashBytes), nil
}
