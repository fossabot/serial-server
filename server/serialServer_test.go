package server

import (
	"aded175/serial-server/conversation"
	"bytes"
	"testing"
)

func TestSerialServer_send(t *testing.T) {
	var stub bytes.Buffer
	s := &serialServer{device: &stub}
	cmd, dat := [4]uint8{1, 2, 3, 4}, []uint8{'t', 'e', 's', 't'}
	expected := []uint8{1, 2, 3, 4, 't', 'e', 's', 't'}
	s.send(cmd, dat...)
	if !bytes.Equal(stub.Bytes(), expected) {
		t.Error("actually:", stub.Bytes(), "expected:", expected)
	}
}

func TestSerialServer_Ping(t *testing.T) {
	var stub bytes.Buffer
	s := &serialServer{device: &stub}
	expected := conversation.PING[:]
	s.Ping()
	if !bytes.Equal(stub.Bytes(), expected) {
		t.Error("actually:", stub.Bytes(), "expected:", expected)
	}
}

func TestSerialServer_Task(t *testing.T) {
	var stub bytes.Buffer
	s := &serialServer{device: &stub}
	data := []uint8{'t', 'e', 's', 't'}
	expected := append(conversation.TASK[:], []uint8{'t', 'e', 's', 't'}...)
	s.Task(data)
	if !bytes.Equal(stub.Bytes(), expected) {
		t.Error("actually:", stub.Bytes(), "expected:", expected)
	}
}
