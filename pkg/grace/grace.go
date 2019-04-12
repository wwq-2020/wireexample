package grace

import (
	"os"
	"os/signal"
	"syscall"
)

func OnExit(cleanup func()) {
	go onExit(cleanup)

}

func onExit(cleanup func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	cleanup()
}
