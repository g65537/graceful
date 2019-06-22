package graceful

import (
	"syscall"
	"time"
)

type option struct {
	restartSignals    []syscall.Signal
	stopSignals       []syscall.Signal
	healthChkInterval time.Duration
	stopTimeout       time.Duration
}

type Option func(o *option)

// WithRestartSignals set reload signals, otherwise, default ones are used
func WithRestartSignals(sigs []syscall.Signal) Option {
	return func(o *option) {
		o.restartSignals = sigs
	}
}

// WithStopSignals set stop signals, otherwise, default ones are used
func WithStopSignals(sigs []syscall.Signal) Option {
	return func(o *option) {
		o.stopSignals = sigs
	}
}

// WithHealthChkInterval set watch interval for worker checking master process state
func WithHealthChkInterval(timeout time.Duration) Option {
	return func(o *option) {
		o.healthChkInterval = timeout
	}
}

// WithStopTimeout set stop timeout for graceful shutdown
//  if timeout occurs, running connections will be discard violently.
func WithStopTimeout(timeout time.Duration) Option {
	return func(o *option) {
		o.stopTimeout = timeout
	}
}
