package file

// Reader is a file reader
type Reader interface {
	Read(path string) (*File, error)
}
