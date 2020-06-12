package server

import (
	"io"
)

type ServerOption func(s *serialServer)

func Options(opts ...ServerOption) ServerOption {
	return func(s *serialServer) {
		for _, opt := range opts {
			opt(s)
		}
	}
}

func Device(reader io.ReadWriter) ServerOption {
	return func(s *serialServer) {
		s.device = reader
	}
}

func WithLogger(logger Logger) ServerOption {
	return func(s *serialServer) {
		s.logger = logger
	}
}
