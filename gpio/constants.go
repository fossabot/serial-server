package gpio

// Address offset: 0x00
const (
	MODE_INPUT  = 0b00 // Input (reset state)
	MODE_OUTPUT = 0b01 // General purpose output mode
	MODE_ALTFUN = 0b10 // Alternate function mode
	MODE_ANALOG = 0b11 // Analog mode
)

// Address offset: 0x04
const (
	OTYPE_PUSHPULL  = 0b00 // Output push-pull (reset state)
	OTYPE_OPENDRAIN = 0b01 // Output open-drain
)

// Address offset: 0x08
const (
	OSPEED_LOW      = 0b00 // Low speed
	OSPEED_MEDIUM   = 0b01 // Medium speed
	OSPEED_HIGH     = 0b10 // High speed
	OSPEED_VERYHIGH = 0b11 // Very high speed
)

// Address offset: 0x0C
const (
	PUPD_NOPULLUP = 0b00 // No pull-up, pull-down
	PUPD_PULLUP   = 0b01 // Pull-up
	PUPD_PULLDOWN = 0b10 // Pull-down
	PUPD_RESERVED = 0b11 // Reserved
)

// Address offset: 0x18
const (
	BSR_NOACTION = 0b00 // No action on the corresponding ODRx bit
	BSR_SET      = 0b01 // Resets the corresponding ODRx bit
)
