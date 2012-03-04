package pubsub

import (
	"fmt"
)

// Message describes a pubsub message.
type Message struct {
	channel string
	data interface{}
}

// String returns a string representation of the message.
func (m *Message) String() string {
	return fmt.Sprintf("%s: %v\n", m.channel, m.data)
}
