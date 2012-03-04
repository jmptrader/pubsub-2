package pubsub

import (
	"fmt"
)

// Message describes a pubsub message.
type Message struct {
	Channel string
	Data interface{}
}

// String returns a string representation of the message.
func (m *Message) String() string {
	return fmt.Sprintf("%s: %v\n", m.Channel, m.Data)
}
