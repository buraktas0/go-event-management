package event

import "sync"

type EventManager struct {
	mutex     sync.Mutex
	listeners map[EventType][]Listener
}

func NewEventManager() *EventManager {
	return &EventManager{listeners: make(map[EventType][]Listener)}
}

func (e *EventManager) AddListener(t EventType, handler func(EventPayload)) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	l := Listener{Channel: make(chan EventPayload), Handler: handler}
	e.listeners[t] = append(e.listeners[t], l)
	go l.Notify(l.Channel, handler)
}

func (e *EventManager) Trigger(t EventType, p EventPayload) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	listeners := e.listeners[t]
	for _, l := range listeners {
		l.Channel <- p
	}
}
