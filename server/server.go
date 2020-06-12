package server

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"hash/crc32"
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

func (s serialServer) checksum() bool {
	return crc32.ChecksumIEEE(s.data[4:]) == binary.BigEndian.Uint32(s.data[:4])
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
			s.data = s.data[:len(s.data)-4]
			if !s.checksum() {
				s.logger.Println("Got corrupted data.")
			}
		}
		fmt.Printf("%x\n", char)
	}
	fmt.Printf("%x, %x, %s\n", s.data[:4], s.data[4:], s.data[4:])
}
