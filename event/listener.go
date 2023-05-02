package event

type Listener struct {
	Channel chan EventPayload
	Handler interface{}
}

func (l *Listener) Notify(ch chan EventPayload, handler func(EventPayload)) {
	for p := range ch {
		data := p
		go func() {
			handler(data)
		}()
	}
}
