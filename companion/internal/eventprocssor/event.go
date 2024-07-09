package eventprocssor

import (
	"companion/internal"
	"companion/internal/config"
)

func (p *processor) processEvent(event internal.StreamEvent) {

	c := config.Get()

	var events []config.Event

	switch event.EventType {
	case internal.EventTypeDono:
		events = c.Events.Dono
	case internal.EventTypeBits:
		events = c.Events.Bits
	case internal.EventTypeGift:
		events = c.Events.Gifts
	case internal.EventTypeT1Sub:
		events = c.Events.Resubt1
	case internal.EventTypeT2Sub:
		events = c.Events.Resubt2
	case internal.EventTypeT3Sub:
		events = c.Events.Resubt3
	}

	ev := findMatch(event.Amount, events)
	if ev != nil {
		squirt(*ev)
	}
}

func findMatch(amount float64, events []config.Event) *config.Event {

	var bestEvent *config.Event

	for _, event := range events {
		if event.Match == "exact" && event.Amount == amount {
			return &event
		}

		if event.Match == "minimum" && event.Amount <= amount {
			if bestEvent == nil || event.Amount > bestEvent.Amount {
				bestEvent = &event
			}
		}
	}

	return bestEvent
}
