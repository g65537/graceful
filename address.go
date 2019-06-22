package graceful

// address defines addr as well as its network type
type address struct {
	addr    string // ip:port, unix path
	network string // tcp, unix
}
