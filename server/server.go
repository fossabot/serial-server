package server

import (
	"bufio"
	"fmt"
	"io"
	"ivan/serial-server/internal/ring"
	"ivan/serial-server/protocol"
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
	buf    *ring.Ring
}

func New(opts ...ServerOption) (ss *serialServer) {
	ss = &serialServer{
		data:   make([]byte, 0, 16),
		logger: &noopLogger{},
		buf:    ring.New(),
	}
	for _, opt := range opts {
		opt(ss)
	}
	return
}

func (s *serialServer) Serve() {
	deviceReader := bufio.NewReader(s.device)
	for {
		char, err := deviceReader.ReadByte()
		if err != nil {
			s.logger.Println(fmt.Errorf("Can not read byte: %w", err))
			continue
		}
		if err == io.EOF {
			break
		}
		s.buf.Add(char)
		if s.buf.Match(protocol.START) {
			s.logger.Println("Got start")
		}
	}
}
