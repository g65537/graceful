package graceful

import (
	"net"
	"net/http"
)

type server struct {
	*http.Server
	listener net.Listener
}
