package pubsub

import (
	. "launchpad.net/gocheck"
	"testing"
	"time"
)

// Hook up gocheck into the gotest runner.
func Test(t *testing.T) {
	TestingT(t)
}

type S struct{}

func init() {
	Suite(&S{})
}

func (s *S) TestBrokerSubscription(c *C) {
	b := NewBroker(10)
	defer b.Close()

	sub := b.Subscription(func(m *Message){}, "chan1", "chan2")
	sub2 := b.Subscription(func(m *Message){}, "chan2")
	time.Sleep(time.Millisecond)
	c.Check(sub.channels["chan1"], Equals, struct{}{})
	c.Check(sub.channels["chan2"], Equals, struct{}{})
	c.Check(len(b.chansubs["chan1"]), Equals, 1)
	c.Check(len(b.chansubs["chan2"]), Equals, 2)
	sub.Subscribe("chan3", "chan4")
	time.Sleep(time.Millisecond)
	c.Check(sub.channels["chan3"], Equals, struct{}{})
	c.Check(sub.channels["chan4"], Equals, struct{}{})
	c.Check(len(b.chansubs["chan3"]), Equals, 1)
	c.Check(len(b.chansubs["chan4"]), Equals, 1)
	sub.Unsubscribe("chan2", "chan3")
	time.Sleep(time.Millisecond)
	_, exists2 := sub.channels["chan2"]
	_, exists3 := sub.channels["chan3"]
	c.Check(exists2, Equals, false)
	c.Check(exists3, Equals, false)
	c.Check(len(b.chansubs["chan2"]), Equals, 1)
	c.Check(len(b.chansubs["chan3"]), Equals, 0)
	sub.Unsubscribe()
	sub2.Unsubscribe()
	time.Sleep(time.Millisecond)
	c.Check(len(sub.channels), Equals, 0)
	c.Check(len(sub2.channels), Equals, 0)
	c.Check(len(b.chansubs["chan1"]), Equals, 0)
	c.Check(len(b.chansubs["chan2"]), Equals, 0)
	c.Check(len(b.chansubs["chan3"]), Equals, 0)
	c.Check(len(b.chansubs["chan4"]), Equals, 0)
}

func (s *S) TestBrokerPublish(c *C) {
	messages1 := []*Message{}
	msgHdlr1 := func(msg *Message) {
		messages1 = append(messages1, msg)
	}

	messages2 := []*Message{}
	msgHdlr2 := func(msg *Message) {
		messages2 = append(messages2, msg)
	}

	b := NewBroker(10)
	defer b.Close()

	sub1 := b.Subscription(msgHdlr1, "chan1", "chan2")
	defer sub1.Unsubscribe()

	sub2 := b.Subscription(msgHdlr2, "chan1")
	defer sub2.Unsubscribe()

	b.Publish("chan1", "mymsg1")
	b.Publish("chan2", "mymsg2")
	sub2.Unsubscribe("chan1")
	b.Publish("chan1", "mymsg3")

	time.Sleep(time.Millisecond)
	c.Assert(len(messages1), Equals, 3)
	c.Check(messages1[0].Channel, Equals, "chan1")
	c.Check(messages1[0].Data, Equals, "mymsg1")
	c.Check(messages1[1].Channel, Equals, "chan2")
	c.Check(messages1[1].Data, Equals, "mymsg2")
	c.Check(messages1[2].Channel, Equals, "chan1")
	c.Check(messages1[2].Data, Equals, "mymsg3")
	c.Assert(len(messages2), Equals, 1)
	c.Check(messages2[0].Channel, Equals, "chan1")
	c.Check(messages2[0].Data, Equals, "mymsg1")

}
