package main

import (
	"errors"
	"ivan/serial-server/server"
	"log"
	"os"

	"github.com/pkg/term"
)

var (
	ErrOpenDevice = errors.New("can not open tty")
)

func main() {
	logger := log.New(os.Stderr, "serial-server ", log.LstdFlags)

	tty, err := term.Open("/dev/pts/1", term.Speed(115200))
	if err != nil {
		logger.Fatalln(ErrOpenDevice)
	}
	defer tty.Close()

	defaultOptions := server.Options(
		server.Device(tty),
		server.WithLogger(logger),
	)
	server := server.New(defaultOptions)

	server.ListenAndServe()
}
