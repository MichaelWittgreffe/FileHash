package hashprocess

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
)

//HashProcessor implements HashProcessor
type HashProcessor struct {
	hasher hash.Hash
}

// NewHashProcessor creates a new instance of a HashProcessor
func NewHashProcessor(hashFuncType string) *HashProcessor {
	var hasher hash.Hash

	switch {
	case hashFuncType == "md5":
		hasher = md5.New()
	case hashFuncType == "sha1":
		hasher = sha1.New()
	case hashFuncType == "sha256":
		hasher = sha256.New()
	case hashFuncType == "sha512":
		hasher = sha512.New()
	default:
		return nil
	}

	return &HashProcessor{hasher: hasher}
}

//Process performs a hash function on the given filepath and returns the hash
func (p *HashProcessor) Process(filepath string) (string, error) {
	if len(filepath) == 0 {
		return "", errors.New("Filepath to Hash Function Is Empty")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("Unable To Open File %s: %s", filepath, err.Error())
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
