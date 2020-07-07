package gpio

type Idr uint32

const IDR_RESET uint32 = 0

func (r Idr) Word() uint16 {
	return uint16(r)
}
