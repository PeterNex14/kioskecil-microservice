package system

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// WaitExitSignal blocks until it receives an interruption signal (SIGINT or SIGTERM).
// It returns the signal received.
func WaitExitSignal() os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	sig := <-stop
	slog.Info("Signal received", "signal", sig.String())
	return sig
}
