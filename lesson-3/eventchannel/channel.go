package eventchannel

import "sync"

type EventHandler struct {
	mu         *sync.Mutex
	channelMap map[string]chan []byte
}

func (h *EventHandler) GetChannel(key string) chan []byte {
	h.mu.Lock()
	ch, ok := h.channelMap[key]
	if !ok {
		ch = make(chan []byte)
		h.channelMap[key] = ch
	}
	h.mu.Unlock()
	return ch
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		mu:         &sync.Mutex{},
		channelMap: make(map[string]chan []byte),
	}
}
