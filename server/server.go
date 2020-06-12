package server

import (
	"bufio"
	"fmt"
	"io"
	"ivan/serial-server/internal/ring"
	"ivan/serial-server/protocol"
)

type serialServer struct {
	device  io.ReadWriter
	logger  Logger
	cmd     *ring.Ring
	session bool
	data    chan byte
}

func New(opts ...ServerOption) (ss *serialServer) {
	ss = &serialServer{
		logger: &noopLogger{},
		cmd:    ring.New(),
		data:   make(chan byte),
	}
	for _, opt := range opts {
		opt(ss)
	}
	return
}

func (s serialServer) send(command [4]uint8, data ...uint8) {
	if _, err := s.device.Write(append(command[:], data...)); err != nil {
		s.logger.Println(fmt.Errorf("can not send. command: %x. data: %q", command, data))
	}
}

func (s serialServer) Ping() {
	s.send(protocol.PING)
}

func (s serialServer) Task(data []uint8) {
	s.send(protocol.TASK, data...)
}

func (s serialServer) ListenAndServe() {
	go s.dispatchData()
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
		s.dispatchCommand()
		if s.session {
			go func() {
				s.data <- char
			}()
		}
	}
}

func (s *serialServer) dispatchCommand() {
	if s.cmd.Match(protocol.PONG) {
	}
	if !s.session && s.cmd.Match(protocol.START) {
		s.session = true
	}
	if s.session && s.cmd.Match(protocol.END) {
		s.session = false
	}
}

func (s serialServer) dispatchData() {
	for {
		c := <-s.data
		fmt.Printf("%x: %c\n", c, c)
	}
}
