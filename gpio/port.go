package gpio

type port struct {
	baseAddress uint32
	moder       Moder
	otyper      Otyper
	ospeedr     Ospeedr
	pupdr       Pupdr
	idr         Idr
	odr         Odr
	bssr        Bsrr
	lckr        Lckr
	afrl        Afrl
	afrh        Afrh
}

const (
	BASE_ADDRESS_PORT_A uint32 = 0x4002_0000
)
