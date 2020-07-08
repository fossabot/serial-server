package gpio

type port struct {
	baseAddress uint32
	Moder
	Otyper
	Ospeedr
	Pupdr
	Idr
	Odr
	Bsrr
	Lckr
	Afrl
	Afrh
}

const (
	BASE_ADDRESS_PORT_A uint32 = 0x4002_0000
)

func NewPortA() *port {
	return &port{
		baseAddress: BASE_ADDRESS_PORT_A,
		Moder:       Moder(MODER_RESET_A),
		Otyper:      Otyper(OTYPER_RESET),
		Ospeedr:     Ospeedr(OSPEEDR_RESET_A),
		Pupdr:       Pupdr(PUPDR_RESET_A),
		Idr:         Idr(IDR_RESET),
		Odr:         Odr(ODR_RESET),
		Bsrr:        Bsrr(BSRR_RESET),
		Lckr:        Lckr(LCKR_RESET),
		Afrl:        Afrl(AFRL_RESET),
		Afrh:        Afrh(AFRH_RESET),
	}
}
