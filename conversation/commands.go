package conversation

const CMD_SIZE uint8 = 4

var (
	PING  = [CMD_SIZE]uint8{0, 0, 0, 0}
	PONG  = [CMD_SIZE]uint8{1, 1, 1, 1}
	TASK  = [CMD_SIZE]uint8{2, 2, 2, 2}
	START = [CMD_SIZE]uint8{3, 3, 3, 3}
	END   = [CMD_SIZE]uint8{4, 4, 4, 4}
	STOP  = [CMD_SIZE]uint8{5, 5, 5, 5}
)
