package hashprocess

//HashProcessor is implemented by anything processing a hash function
type HashProcessor interface {
	Process() (string, error)
}
