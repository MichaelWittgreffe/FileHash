package hashprocess

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strings"
)

//GetHashProcessor returns a hash function for the given hash type
func GetHashProcessor(filepath, hashFuncType string) *HashProcessor {
	var hasher hash.Hash

	if len(hashFuncType) == 0 {
		return nil
	}

	hashFuncType = strings.ToLower(hashFuncType)

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

	return &HashProcessor{filepath: filepath, hasher: hasher}
}
