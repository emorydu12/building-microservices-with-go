package deadline

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timeout waiting for function to finish")

type Deadline struct {
	timeout time.Duration
}

func New(timeout time.Duration) *Deadline {
	return &Deadline{
		timeout: timeout,
	}
}

type WorkFunc func(<-chan struct{}) error

func (d *Deadline) Run(workfn WorkFunc) error {
	result := make(chan error, 1)
	stopper := make(chan struct{})

	go func() {
		result <- workfn(stopper)
	}()

	timer := time.NewTimer(d.timeout)
	select {
	case err := <-result:
		timer.Stop()
		return err
	case <-timer.C:
		close(stopper)
		return ErrTimeout
	}
}
