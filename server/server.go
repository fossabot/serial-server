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

type serialServer struct {
	device io.ReadWriter
	logger Logger

	data    []byte
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
		s.logger.Printf("Can not send. Command: %x. Data: %q.", command, data)
	}
}

func (s serialServer) Ping() {
	s.send(protocol.PING)
}

func (s serialServer) Test(data []uint8) {
	s.send(protocol.TEST, data...)
}

func (s serialServer) Serve() {
	s.CommandListener()
}

func (s serialServer) CommandListener() {
	deviceReader := bufio.NewReader(s.device)
	for {
		char, err := deviceReader.ReadByte()
		if err == io.EOF {
			s.logger.Println(fmt.Errorf("got: %w", io.EOF))
			break
		}
		if char == 0 && err != nil {
			s.logger.Println(fmt.Errorf("got %x: %w", char, err))
			break
		}
		if err != nil {
			s.logger.Println(fmt.Errorf("skip: %w", err))
			continue
		}
		s.cmd.Add(char)
	}
}
