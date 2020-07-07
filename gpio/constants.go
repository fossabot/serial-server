package gpio

import "errors"

type moderState uint8

const (
	MODE_INPUT  moderState = 0b00 // Input (reset state)
	MODE_OUTPUT moderState = 0b01 // General purpose output mode
	MODE_ALTFUN moderState = 0b10 // Alternate function mode
	MODE_ANALOG moderState = 0b11 // Analog mode
)

const (
	MODE_RESET_A uint32 = 0xA800_0000
	MODE_RESET_B uint32 = 0x0280
	MODE_RESET   uint32 = 0
)

const MODE_PIN_SIZE uint8 = 2

type otyperState uint8

const (
	OTYPE_PUSHPULL  otyperState = 0b00 // Output push-pull (reset state)
	OTYPE_OPENDRAIN otyperState = 0b01 // Output open-drain
)

const OTYPE_RESET uint32 = 0

const (
	OSPEED_LOW      = 0b00 // Low speed
	OSPEED_MEDIUM   = 0b01 // Medium speed
	OSPEED_HIGH     = 0b10 // High speed
	OSPEED_VERYHIGH = 0b11 // Very high speed
)

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
