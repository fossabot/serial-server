package config

import (
	"flag"
)

type config struct {
	Device   string
	BaudRate int
}

func FromFlags() *config {
	c := &config{}
	flag.StringVar(&c.Device, "device", "", "path to serial device")
	flag.IntVar(&c.BaudRate, "rate", 115200, "connection baud rate")
	flag.Parse()
	return c
}
