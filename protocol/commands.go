package protocol

var (
	END   = [4]uint8{0, 0, 0, 0}
	START = [4]uint8{1, 1, 1, 1}
	PING  = [4]uint8{2, 2, 2, 2}
	PONG  = [4]uint8{3, 3, 3, 3}
	TEST  = [4]uint8{4, 4, 4, 4}
)
