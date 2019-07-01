package hashprocess

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

//SupportedTypes is a simple array of lower-case string hash function names (eg: "md5")
var SupportedTypes = [4]string{"md5", "sha1", "sha256", "sha512"}

//GetHashProcessor returns a hash function for the given hash type
func GetHashProcessor(filepath, hashFuncType string) *HashProcessor {
	switch {
	case hashFuncType == "md5":
		return &HashProcessor{filepath: filepath, hasher: md5.New()}
	case hashFuncType == "sha1":
		return &HashProcessor{filepath: filepath, hasher: sha1.New()}
	case hashFuncType == "sha256":
		return &HashProcessor{filepath: filepath, hasher: sha256.New()}
	case hashFuncType == "sha512":
		return &HashProcessor{filepath: filepath, hasher: sha512.New()}
	default:
		return nil
	}
}
