package gpio

import "testing"

func TestModer_Set(t *testing.T) {
	var testModer = Moder(MODE_RESET)
	testModer.Set(3, MODE_OUTPUT)
	if uint32(testModer) != 0b01_00_00_00 {
		t.Error("set incorrect")
	}
}

func TestModer_Test(t *testing.T) {
	var moder = Moder(0b11_00_00)
	if !moder.Test(2, MODE_ANALOG) {
		t.Error("3rd byte does not set")
	}
}

func TestModer_Test_False(t *testing.T) {
	var moder = Moder(0b11_00_00)
	if moder.Test(3, MODE_ANALOG) {
		t.Error("4rd byte set")
	}
}
