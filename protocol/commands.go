package protocol

var (
	PING  = [4]uint8{0, 0, 0, 0}
	PONG  = [4]uint8{1, 1, 1, 1}
	TASK  = [4]uint8{2, 2, 2, 2}
	START = [4]uint8{3, 3, 3, 3}
	END   = [4]uint8{4, 4, 4, 4}
)
