package hashprocess

//MD5Processor implements HashProcessor
type MD5Processor struct {
	filepath string
}

//Process performs a hash function on the given filepath and returns the hash
func (p *MD5Processor) Process() (string, error) {
	return "md5 hashstring", nil
}
