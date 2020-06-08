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
	device  io.Reader
	data    []byte
	logger  Logger
	cmd     *ring.Ring
	session bool
}

func New(opts ...ServerOption) (ss *serialServer) {
	ss = &serialServer{
		data:   make([]byte, 0, 16),
		logger: &noopLogger{},
		cmd:    ring.New(),
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
		if err == io.EOF {
			s.logger.Println(fmt.Errorf("Got EOF. Stoping serve."))
			break
		}
		if err != nil {
			s.logger.Println(fmt.Errorf("Can not read byte: %w", err))
			continue
		}
		if s.session {
			s.data = append(s.data, char)
		}
		s.cmd.Add(char)
		if !s.session && s.cmd.Match(protocol.START) {
			s.session = true
		}
		if s.session && s.cmd.Match(protocol.END) {
			s.session = false
		}
	}
}
