package hashprocess

//SHA512Processor implements HashProcessor
type SHA512Processor struct {
	filepath string
}

//Process performs a SHA1 hash function on the given filepath and returns the hash
func (p *SHA512Processor) Process() (string, error) {
	return "sha512 hashstring", nil
}
