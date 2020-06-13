package ring

import (
	"aded175/serial-server/protocol"
	"bytes"
)

type Ring struct {
	buf  [protocol.CMD_SIZE]byte
	head uint8
	tail uint8
	max  uint8
	full bool
}

func New() *Ring {
	return &Ring{
		max: protocol.CMD_SIZE,
	}
}

func (r *Ring) Add(data byte) {
	r.buf[r.head] = data
	r.moveForward()
}

func (r *Ring) moveForward() {
	if r.full {
		r.tail = (r.tail + 1) % r.max
	}
	r.head = (r.head + 1) % r.max
	r.full = r.head == r.tail
}

func (r *Ring) Match(data [protocol.CMD_SIZE]byte) bool {
	return bytes.Equal(r.buf[:], data[:])
}
