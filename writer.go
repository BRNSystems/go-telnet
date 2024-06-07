package telnet

type Writer interface {
	Write([]byte) (int, error)
	RawWrite([]byte) (int, error)
}
