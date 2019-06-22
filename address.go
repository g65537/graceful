package graceful

// address defines addr as well as its network type
type address struct {
	network string // tcp, unix
	addr    string // ip:port, unix path
}
