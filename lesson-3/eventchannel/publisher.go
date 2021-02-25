package eventchannel

type Publisher struct {
	chanKeys     []string
	eventHandler *EventHandler
}

func NewPublisher(eventHandler *EventHandler, chanKeys []string) *Publisher {
	return &Publisher{
		eventHandler: eventHandler,
		chanKeys:     chanKeys,
	}
}

func (pub *Publisher) Publish(data []byte) {
	for _, key := range pub.chanKeys {
		pub.eventHandler.GetChannel(key) <- data
	}
}
