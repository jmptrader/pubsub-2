// The pubsub package implements a publish-subscribe messaging pattern.
package pubsub

// Subscription is structure that holds data for broker subscriptions.
type Subscription struct {
	broker *Broker
	msgHdlr func(msg *Message)
	channels map[string]struct{}
}

func newSubscription(broker *Broker, msgHdlr func(msg *Message), channels []string) *Subscription {
	s := &Subscription{
		broker: broker,
		msgHdlr: msgHdlr,
	channels: make(map[string]struct{}),
	}

	s.Subscribe(channels...)

	return s
}

// Subscribe queues a subcribe to given channels.
func (s *Subscription) Subscribe(channels ...string) {
	if len(channels) < 1 {
		return
	}

	s.broker.dataChan <- &envData{false, &envSubscription{true, s, channels}, nil}
}

// Unsubscribe queues a unsubscribe from the given channels.
// If no channels are given, the Subscription will be unsubscribed from all currently subscribed
// channels.
func (s *Subscription) Unsubscribe(channels ...string) {
	s.broker.dataChan <- &envData{false, &envSubscription{false, s, channels}, nil}
}
