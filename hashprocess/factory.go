package hashprocess

//SupportedTypes is a simple array of lower-case string hash function names (eg: "md5")
var SupportedTypes = [4]string{"md5", "sha1", "sha256", "sha512"}

//GetHashFunction returns a hash function for the given hash type
func GetHashFunction(filepath, hashFuncType string) *HashProcessor {
	var result HashProcessor

	switch {
	case hashFuncType == "md5":
		result = &MD5Processor{filepath: filepath}
	case hashFuncType == "sha1":
		result = &SHA1Processor{filepath: filepath}
	case hashFuncType == "sha256":
		result = &SHA256Processor{filepath: filepath}
	case hashFuncType == "sha512":
		result = &SHA512Processor{filepath: filepath}
	}

	return &result
}
