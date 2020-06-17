package conversation

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
)

type Pin uint64

type PinState uint8

const (
	PIN_A Pin = 0x1000000000
	PIN_B Pin = 0x1100000000
)

const (
	LOW  PinState = 0
	HIGH PinState = 1
)

type Pair struct {
	Pin
	PinState
}

type Conversation struct {
	Pairs []Pair
}

func (c Conversation) Marshal() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 256))
	for _, p := range c.Pairs {
		err := binary.Write(b, binary.BigEndian, p.Pin)
		if err != nil {
			return nil, err
		}
		err = binary.Write(b, binary.BigEndian, p.PinState)
		if err != nil {
			return nil, err
		}
	}
	err := binary.Write(b, binary.BigEndian, crc32.ChecksumIEEE(b.Bytes()))
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
