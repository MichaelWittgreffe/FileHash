package hashprocess

import (
	"encoding/hex"
	"errors"
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
	var err error
	result := ""

	if len(p.filepath) > 0 {
		if file, err := os.Open(p.filepath); err == nil {
			if _, err = io.Copy(p.hasher, file); err == nil {
				if hashBytes := p.hasher.Sum(nil); len(hashBytes) > 0 {
					result = hex.EncodeToString(hashBytes)
				}
			}
		}
	} else {
		err = errors.New("Filepath To Hash Function Is Empty")
	}

	return result, err
}
