package session

import "testing"

func TestSessionManager_Toggle(t *testing.T) {
	s := New()
	if s.Active() {
		t.Errorf("session must be inactive, actually %t", s.Active())
	}
	s.Toggle()
	if !s.Active() {
		t.Errorf("session must be active, actually %t", s.Active())
	}
}
