package session

import "fmt"

type SessionManager interface {
	// Toggle session state
	Toggle()
	// Check is session active or not
	Active() bool
	// Add a byte of data
	Add(uint8)
	Drain()
}

type sessionManager struct {
	active bool
	data   chan uint8
}

func New() SessionManager {
	return &sessionManager{
		data: make(chan uint8),
	}
}

func (sm sessionManager) Close() {
	close(sm.data)
}

func (sm *sessionManager) Toggle() {
	sm.active = !sm.active
}

func (sm sessionManager) Active() bool {
	return sm.active
}

func (sm sessionManager) Add(char uint8) {
	if sm.Active() {
		sm.data <- char
	}
}

func (sm sessionManager) Drain() {
	for {
		c := <-sm.data
		fmt.Printf("%x: %c\n", c, c)
	}
}
