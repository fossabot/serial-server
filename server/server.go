package server

import (
	"encoding/binary"
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
	device  io.ReadWriter
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

func (s serialServer) send(command [4]uint8, data ...uint8) {
	if _, err := s.device.Write(append(command[:], data...)); err != nil {
		s.logger.Println("Can not send", data)
	}
}

func (s serialServer) Ping() {
	s.send(protocol.PING)
}

func (s serialServer) Test(data []uint8) {
	s.send(protocol.TEST, data...)
}

func (s serialServer) Serve() {
	s.Test([]uint8("test!"))
}
