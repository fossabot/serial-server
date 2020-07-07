package gpio

type pupdrState uint8

const (
	PUPD_NOPULLUP pupdrState = 0b00 // No pull-up, pull-down
	PUPD_PULLUP   pupdrState = 0b01 // Pull-up
	PUPD_PULLDOWN pupdrState = 0b10 // Pull-down
	PUPD_RESERVED pupdrState = 0b11 // Reserved
)

const (
	PUPDR_RESET_A uint32 = 0x6400_0000
	PUPDR_RESET_B uint32 = 0x0100
	PUPDR_RESET   uint32 = 0
)

const PUPDR_PIN_SIZE uint8 = 2

type Pupdr uint32
