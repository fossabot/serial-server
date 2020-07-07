package gpio

type port struct {
	baseAddress uint32
	moder       Moder
	otyper      Otyper
	ospeedr     Ospeedr
	pupdr       Pupdr
	idr         uint32
	odr         uint32
	bssr        uint32
	lckr        uint32
	afrl        uint32
	afrh        uint32
}
