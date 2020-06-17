package conversation

import "fmt"

type ConversationManager interface {
	// Toggle session state
	Toggle()
	// Check is session active or not
	Active() bool
	// Add a byte of data
	Add(uint8)
	Drain()
}

type conversationManager struct {
	active bool
	data   chan uint8
}

func New() ConversationManager {
	return &conversationManager{
		data: make(chan uint8),
	}
}

func (sm conversationManager) Close() {
	close(sm.data)
}

func (sm *conversationManager) Toggle() {
	sm.active = !sm.active
}

func (sm conversationManager) Active() bool {
	return sm.active
}

func (sm conversationManager) Add(char uint8) {
	if sm.Active() {
		sm.data <- char
	}
}

func (sm conversationManager) Drain() {
	for {
		c := <-sm.data
		fmt.Printf("%x: %c\n", c, c)
	}
}
