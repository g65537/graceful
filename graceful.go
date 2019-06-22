package graceful

import (
	"net/http"
	"os"
	"syscall"
	"time"
)

type GracefulSvr struct {
	opt      *option
	addrs    []address
	handlers []http.Handler
}

func NewGracefulSvr(opts ...Option) *GracefulSvr {
	option := &option{
		restartSignals:    defaultReloadSignals,
		stopSignals:       defaultStopSignals,
		healthChkInterval: defaultWatchInterval,
		stopTimeout:       defaultStopTimeout,
	}
	for _, opt := range opts {
		opt(option)
	}
	return &GracefulSvr{
		addrs:    make([]address, 0),
		handlers: make([]http.Handler, 0),
		opt:      option,
	}
}

// Register an addr and its corresponding handler
// all (addr, handler) pair will be started with server.Run
func (s *GracefulSvr) Register(addr string, handler http.Handler) {
	s.addrs = append(s.addrs, address{addr, "tcp"})
	s.handlers = append(s.handlers, handler)
}

// RegisterUnix register (addr, handler) on unix socket
func (s *GracefulSvr) RegisterUnix(addr string, handler http.Handler) {
	s.addrs = append(s.addrs, address{addr, "unix"})
	s.handlers = append(s.handlers, handler)
}

// Run runs all register servers
func (s *GracefulSvr) Run() error {
	if len(s.addrs) == 0 {
		return ErrNoServers
	}
	StartedAt = time.Now()
	if IsWorker() {
		worker := &worker{handlers: s.handlers, opt: s.opt, stopCh: make(chan struct{})}
		return worker.run()
	}
	master := &master{addrs: s.addrs, opt: s.opt, workerExit: make(chan error)}
	return master.run()
}

// Reload reload server gracefully
func (s *GracefulSvr) Reload() error {
	ppid := os.Getppid()
	if IsWorker() && ppid != 1 && len(s.opt.restartSignals) > 0 {
		return syscall.Kill(ppid, s.opt.restartSignals[0])
	}

	// Reload called by user from outside usally in user's handler
	// which works on worker, master don't need to handle this
	return nil
}
