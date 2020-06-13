package server

import (
	"io"
)

type serverOption func(s *serialServer)

func Options(opts ...serverOption) serverOption {
	return func(s *serialServer) {
		for _, opt := range opts {
			opt(s)
		}
	}
}

func Device(dev io.ReadWriter) serverOption {
	return func(s *serialServer) {
		s.device = dev
	}
}

func WithLogger(logger Logger) serverOption {
	return func(s *serialServer) {
		s.logger = logger
	}
}
