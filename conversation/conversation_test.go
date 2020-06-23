package conversation

import (
	"fmt"
)

func ExampleConversation_Marshal() {
	conv := Conversation{
		Pairs: []Pair{
			{
				Pin:      0x10,
				PinState: 0x01,
			},
		},
	}
	bytes, _ := conv.Marshal()
	fmt.Printf("%X\n", bytes)
	//Output:
	//0000001001FBE7D5DA
}
