package shared

import "sync"

type Event interface {
	Topic() string
}

type Listener interface {
	Handle(event Event)
}

type EventBus struct {
	listeners map[string][]Listener
	mu        sync.Mutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[string][]Listener),
	}
}

func (bus *EventBus) Subscribe(topic string, listener Listener) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	bus.listeners[topic] = append(bus.listeners[topic], listener)
}

func (bus *EventBus) Publish(event Event) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	if listeners, found := bus.listeners[event.Topic()]; found {
		for _, listener := range listeners {
			go listener.Handle(event)
		}
	}
}
