package eventprocssor

import (
	"companion/internal"
	"companion/internal/config"
	"time"
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
	go p.run()
}

type processor struct {
	config   config.Config
	events   <-chan internal.StreamEvent
	messages <-chan internal.ChatMessage
}

func (p *processor) receive() {
	cooldownEnd := time.Now()
	for {
		select {
		case event := <-p.events:
			if time.Now().Before(cooldownEnd) {
				continue
			}
			p.processEvent(event)
		case message := <-p.messages:
			if time.Now().Before(cooldownEnd) {
				continue
			}
			p.processMessage(message)
		}
		cooldownEnd = time.Now().Add(time.Duration(p.config.Settings.GlobalCooldown))
	}
}
