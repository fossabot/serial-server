package gpio

type Odr uint32

const ODR_RESET uint32 = 0

func (r *Odr) Set(pinOffset, state uint8) error {
	if pinOffset > 15 {
		return ErrPinNotExists
	}
	*r |= Odr(state << pinOffset)
	return nil
}
