package pubsub

type envSubscription struct {
	subscription bool
	sub *Subscription
	channels []string
}

type envData struct {
	closing bool
	subscription *envSubscription
	message *Message
}

//* Broker

// Broker is a message broker which transmits published messages to its subscribers.
type Broker struct {
	// mapping from channels to subscription sets
	chansubs map[string]map[*Subscription]struct{}
	dataChan chan *envData
}

// Return a new Broker with the given message buffer size.
func NewBroker(bufSize int) *Broker {
	b := &Broker{
	chansubs: make(map[string]map[*Subscription]struct{}),
	dataChan: make(chan *envData, bufSize),
	}

	go b.backend()
	return b
}

// Publish queues a message to be sent to the given channels with the given data.
func (b *Broker) Publish(channel string, data interface{}) {
	b.dataChan <- &envData{false, nil, &Message{channel, data}}
}

// Subscription returns a new Subscription with the given message handler function and
// queues it to subscribe to the given channels.
func (b *Broker) Subscription(msgHdlr func(msg *Message), channels ...string) *Subscription {
	return newSubscription(b, msgHdlr, channels)
}

// Close queues a closing to the broker.
func (b *Broker) Close() {
	b.dataChan <- &envData{true, nil, nil}
}

func (b *Broker) backend() {
	for data := range b.dataChan {
		switch {
		case data.closing:
			close(b.dataChan)
		case data.subscription != nil:
			es := data.subscription
			if es.subscription {
				// subscribe
				for _, channel := range es.channels {
					if _, exists := b.chansubs[channel]; !exists {
						// create new channel
						b.chansubs[channel] = make(map[*Subscription]struct{})
					}

					b.chansubs[channel][es.sub] = struct{}{}
					es.sub.channels[channel] = struct{}{}
				}
			} else {
				// unsubscribe
				var channels []string

				if len(es.channels) > 0 {
					channels = es.channels
				} else {
					// unsubscribe from all channels
					for channel, _ := range es.sub.channels {
						channels = append(channels, channel)
					}
					es.sub.channels = map[string]struct{}{}
				}
				
				for _, channel := range channels {
					delete(b.chansubs[channel], es.sub)
					if len(b.chansubs[channel]) == 0 {
						// delete empty channel
						delete(b.chansubs, channel)
					}
					
					delete(es.sub.channels, channel)
				}
			}
		case data.message != nil:
			msg := data.message
			subset, exists := b.chansubs[msg.channel]
			if exists {
				for sub, _ := range subset {
					sub.msgHdlr(msg)
				}
			}
		}
	}
}