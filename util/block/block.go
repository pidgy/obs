package block

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

// For is a utility function to block on the calling thread for a duration d.
func For(d time.Duration) {
	sigq := make(chan os.Signal, 1)
	signal.Notify(sigq, syscall.SIGINT, syscall.SIGTERM)

	okq := make(chan bool)
	time.AfterFunc(d, func() { close(okq) })

	select {
	case <-okq:
	case <-sigq:
	}
}
