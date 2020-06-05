package ring

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Add 2 values", func(t *testing.T) {
		ring := New()
		ring.Add(byte('a'))
		ring.Add(byte('b'))

		if ring.full != false {
			t.Error("Not full expected, actually full")
		}
		if ring.head != 2 {
			t.Errorf("Head at 2 epected, actually %d", ring.head)
		}
		if ring.tail != 0 {
			t.Errorf("Tail at 0 epected, actually %d", ring.tail)
		}
	})

	t.Run("Add 4 values", func(t *testing.T) {
		ring := New()
		ring.Add(byte('a'))
		ring.Add(byte('b'))
		ring.Add(byte('c'))
		ring.Add(byte('d'))

		if ring.full != true {
			t.Error("Full expected, actually not full")
		}
		if ring.head != 0 {
			t.Errorf("Head at 0 epected, actually %d", ring.head)
		}
		if ring.tail != 0 {
			t.Errorf("Tail at 0 epected, actually %d", ring.tail)
		}
	})

	t.Run("Add 5 values", func(t *testing.T) {
		ring := New()
		ring.Add(byte('a'))
		ring.Add(byte('b'))
		ring.Add(byte('c'))
		ring.Add(byte('d'))
		ring.Add(byte('e'))

		if ring.full != true {
			t.Error("Full expected, actually not full")
		}
		if ring.head != 1 {
			t.Errorf("Head at 0 epected, actually %d", ring.head)
		}
		if ring.tail != 1 {
			t.Errorf("Tail at 0 epected, actually %d", ring.tail)
		}
	})
}

func TestMatch(t *testing.T) {
	ring := New()
	ring.Add(byte('a'))
	ring.Add(byte('b'))

	if !ring.Match([4]byte{'a', 'b'}) {
		t.Errorf("Does not match 2 values")
	}
}
