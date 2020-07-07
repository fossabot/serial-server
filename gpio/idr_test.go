package gpio

import "testing"

func TestIdr_Word(t *testing.T) {
	var idr = Idr(0x1_0000)
	if idr.Word() != 0 {
		t.Error("incorrect truncate")
	}
}
