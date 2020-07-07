package gpio

import "errors"

const (
	PUPD_NOPULLUP = 0b00 // No pull-up, pull-down
	PUPD_PULLUP   = 0b01 // Pull-up
	PUPD_PULLDOWN = 0b10 // Pull-down
	PUPD_RESERVED = 0b11 // Reserved
)

const (
	BSR_NOACTION = 0b00 // No action on the corresponding ODRx bit
	BSR_SET      = 0b01 // Resets the corresponding ODRx bit
)

var ErrPinNotExists = errors.New("pin does not exists")
