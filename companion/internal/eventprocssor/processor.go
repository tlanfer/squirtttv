package eventprocssor

import (
	"companion/internal"
	"companion/internal/config"
)

func New(events <-chan internal.StreamEvent, messages <-chan internal.ChatMessage) {
	c := config.Get()
	p := processor{
		config:   c,
		events:   events,
		messages: messages,
	}

	config.Subscribe(func(c config.Config) {
		p.config = c
	})

	go p.receive()
	go run()
}

type processor struct {
	config   config.Config
	events   <-chan internal.StreamEvent
	messages <-chan internal.ChatMessage
}

func (p *processor) receive() {
	for {
		select {
		case event := <-p.events:
			p.processEvent(event)
		case message := <-p.messages:
			p.processMessage(message)
		}
	}
}

func (p *processor) processMessage(msg internal.ChatMessage) {
}
