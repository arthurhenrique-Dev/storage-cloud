package events

type Event interface{}

type Handler func(Event)

type Publisher struct {
	handlers map[string][]Handler
}

func NewPublisher() *Publisher {
	return &Publisher{
		handlers: make(map[string][]Handler),
	}
}

func (p *Publisher) Subscribe(eventName string, handler Handler) {
	p.handlers[eventName] = append(
		p.handlers[eventName],
		handler,
	)
}

func (p *Publisher) Publish(eventName string, event Event) {
	for _, handler := range p.handlers[eventName] {
		go handler(event)
	}
}
