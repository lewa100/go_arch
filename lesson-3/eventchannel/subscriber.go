package eventchannel

type Subscriber struct {
	chanKey      string
	eventHandler *EventHandler
}

func (sub *Subscriber) GetMessage() []byte {
	return <-sub.eventHandler.GetChannel(sub.chanKey)
}

func NewSubscriber(eventHandler *EventHandler, chanKey string) *Subscriber {
	return &Subscriber{
		eventHandler: eventHandler,
		chanKey:      chanKey,
	}
}
