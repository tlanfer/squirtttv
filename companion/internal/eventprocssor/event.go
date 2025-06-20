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
		squirt(ev.Choose, ev.Devices, ev.Pattern)
	}
}
