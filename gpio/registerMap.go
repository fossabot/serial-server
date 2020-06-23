package gpio

type registerMap struct {
	moder   uint32
	otyper  uint32
	ospeedr uint32
	pupdr   uint32
	idr     uint32
	odr     uint32
	bssr    uint32
	lckr    uint32
	afrl    uint32
	afrh    uint32
}
