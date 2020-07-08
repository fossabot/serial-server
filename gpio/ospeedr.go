package gpio

type ospeedrState uint8

const (
	OSPEED_LOW      ospeedrState = 0b00 // Low speed
	OSPEED_MEDIUM   ospeedrState = 0b01 // Medium speed
	OSPEED_HIGH     ospeedrState = 0b10 // High speed
	OSPEED_VERYHIGH ospeedrState = 0b11 // Very high speed
)

const (
	OSPEEDR_RESET_A uint32 = 0x0C00_0000
	OSPEEDR_RESET_B uint32 = 0x00C0
	OSPEEDR_RESET   uint32 = 0
)

const OSPEED_PIN_SIZE uint8 = 2

type Ospeedr uint32
