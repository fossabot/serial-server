package server

import (
	"aded175/serial-server/conversation"
	"aded175/serial-server/internal/ring"
	"aded175/serial-server/protocol"
	"bufio"
	"fmt"
	"io"
)

type serialServer struct {
	device io.ReadWriter
	logger Logger
	cmd    *ring.Ring
	conversation.ConversationManager
}

func New(opts ...serverOption) (ss *serialServer) {
	ss = &serialServer{
		logger:              &noopLogger{},
		cmd:                 ring.New(),
		ConversationManager: conversation.New(),
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
		s.Add(char)
		s.cmd.Add(char)
		if !s.Active() && s.cmd.Match(protocol.START) ||
			s.Active() && s.cmd.Match(protocol.END) {
			s.Toggle()
		}
	}
}
