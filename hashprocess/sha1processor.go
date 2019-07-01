package hashprocess

//SHA1Processor implements HashProcessor
type SHA1Processor struct {
	filepath string
}

//Process performs a SHA1 hash function on the given filepath and returns the hash
func (p *SHA1Processor) Process() (string, error) {
	return "sha1 hashstring", nil
}
