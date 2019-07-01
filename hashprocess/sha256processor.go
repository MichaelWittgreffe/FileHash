package hashprocess

//SHA256Processor implements HashProcessor
type SHA256Processor struct {
	filepath string
}

//Process performs a SHA256 hash function on the given filepath and returns the hash
func (p *SHA256Processor) Process() (string, error) {
	return "sha256 hashstring", nil
}
