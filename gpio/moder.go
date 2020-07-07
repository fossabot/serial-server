package gpio

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
