package gpio

type bsrrState uint8

const (
	BSRR_NOACTION  bsrrState = 0b00 // No action on the corresponding ODRx bit
	BSRR_SET_RESET bsrrState = 0b01 // Set/Reset the corresponding ODRx bit
)

const BSRR_RESET uint32 = 0

type Bsrr uint32

func (r *Bsrr) SetBs(pinOffset uint8, state bsrrState) error {
	if pinOffset > 15 {
		return ErrPinNotExists
	}
	*r |= Bsrr(state << pinOffset)
	return nil
}

func (r *Bsrr) SetBr(pinOffset uint8, state bsrrState) error {
	if pinOffset > 15 {
		return ErrPinNotExists
	}
	*r |= Bsrr(state << (pinOffset + 15))
	return nil
}

func (r Bsrr) TestBs(pinOffset uint8, state bsrrState) bool {
	return (uint32(r)>>(pinOffset))&uint32(state) == uint32(state)
}

func (r Bsrr) TestBr(pinOffset uint8, state bsrrState) bool {
	return (uint32(r)>>(pinOffset+15))&uint32(state) == uint32(state)
}
