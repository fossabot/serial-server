package protocol

type command [4]uint8

func newCommand(a uint8, b uint8, c uint8, d uint8) command {
	return command([4]uint8{a, b, c, d})
}

var (
	END   = newCommand(0, 0, 0, 0)
	START = newCommand(1, 1, 1, 1)
	PING  = newCommand(2, 2, 2, 2)
	PONG  = newCommand(3, 3, 3, 3)
	TEST  = newCommand(4, 4, 4, 4)
)
