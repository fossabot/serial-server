package gpio

type otyperState uint8

const (
	OTYPE_PUSHPULL  otyperState = 0b00 // Output push-pull (reset state)
	OTYPE_OPENDRAIN otyperState = 0b01 // Output open-drain
)

const OTYPER_RESET uint32 = 0

type Otyper uint32

func (r *Otyper) Set(pinOffset uint8, state otyperState) error {
	if pinOffset > 15 {
		return ErrPinNotExists
	}
	*r |= Otyper(state << pinOffset)
	return nil
}

func (r Otyper) Test(pinOffset uint8, state otyperState) bool {
	return (uint32(r)>>pinOffset)&uint32(state) == uint32(state)
}
