package server

import (
	"bufio"
	"fmt"
	"io"
)

type Logger interface {
	Println(v ...interface{})
}

type noopLogger struct{}

func (l noopLogger) Println(v ...interface{}) {}

type serialServer struct {
	device io.Reader
	data   []byte
	logger Logger
}

func New(opts ...ServerOption) (ss *serialServer) {
	ss = &serialServer{
		data:   make([]byte, 0, 16),
		logger: &noopLogger{},
	}
	for _, opt := range opts {
		opt(ss)
	}
	return
}

func (s *serialServer) Serve() {
	deviceReader := bufio.NewReader(s.device)
	for {
		dataByte, err := deviceReader.ReadByte()
		if err != nil {
			s.logger.Println(fmt.Errorf("Can not read byte: %w", err))
		}
		if err == io.EOF {
			break
		}
		s.data = append(s.data, dataByte)
	}
}
