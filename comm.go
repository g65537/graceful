package graceful

import (
	"net/http"
	"os"
	"syscall"
	"time"
)

// constants
const (
	EnvWorker       = "GRACEFUL_WORKER"
	EnvNumFD        = "GRACEFUL_NUMFD"
	EnvOldWorkerPid = "GRACEFUL_OLD_WORKER_PID"
	ValWorker       = "1"
)

var (
	defaultWatchInterval = time.Second
	defaultStopTimeout   = 20 * time.Second
	defaultReloadSignals = []syscall.Signal{syscall.SIGHUP, syscall.SIGUSR1}
	defaultStopSignals   = []syscall.Signal{syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT}

	StartedAt time.Time
)

// ListenAndServe starts server with (addr, handler)
func ListenAndServe(addr string, handler http.Handler) error {
	server := NewGracefulSvr()
	server.Register(addr, handler)
	return server.Run()
}

func IsWorker() bool {
	return os.Getenv(EnvWorker) == ValWorker
}

func IsMaster() bool {
	return !IsWorker()
}
