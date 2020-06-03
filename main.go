package main

import (
	"errors"
	"ivan/serial-server/server"
	"log"
	"os"
)

var (
	ErrOpenDevice = errors.New("can not open tty")
)

func main() {
	logger := log.New(os.Stderr, "serial-server ", log.LstdFlags)

	tty, err := os.Open("/dev/pts/2")
	if err != nil {
		logger.Fatalln(ErrOpenDevice)
	}
	defer tty.Close()

	defaultOptions := server.Options(
		server.Device(tty),
		server.WithLogger(logger),
	)
	server := server.New(defaultOptions)

	server.Serve()
}
