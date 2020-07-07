package gpio

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

type Moder uint32

func (r *Moder) Set(pinOffset uint8, state moderState) error {
	if pinOffset > 31 {
		return ErrPinNotExists
	}
	*r |= Moder(state << (pinOffset * MODE_PIN_SIZE))
	return nil
}

func (r Moder) Test(pinOffset uint8, state moderState) bool {
	return (uint32(r)>>(pinOffset*MODE_PIN_SIZE))&uint32(state) == uint32(state)
}
