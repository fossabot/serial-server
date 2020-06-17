package main

import (
	"aded175/serial-server/cmd/server/config"
	"aded175/serial-server/server"
	"errors"
	"log"
	"os"

	"github.com/pkg/term"
)

var (
	ErrOpenDevice = errors.New("can not open tty")
)

func main() {
	cfg := config.FromFlags()
	logger := log.New(os.Stderr, "serial-server ", log.LstdFlags)

	tty, err := term.Open(cfg.Device, term.Speed(cfg.BaudRate))
	if err != nil {
		logger.Fatalln(ErrOpenDevice)
	}
	defer tty.Close()

	defaultOptions := server.Options(
		server.Device(tty),
		server.WithLogger(logger),
	)
	server := server.New(defaultOptions)

	go server.Drain()
	server.ListenAndServe()
}
