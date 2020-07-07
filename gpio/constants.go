package gpio

import "errors"

const (
	BSR_NOACTION = 0b00 // No action on the corresponding ODRx bit
	BSR_SET      = 0b01 // Resets the corresponding ODRx bit
)

var ErrPinNotExists = errors.New("pin does not exists")
