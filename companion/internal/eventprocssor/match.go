package eventprocssor

import (
	"companion/internal/config"
)

func findMatch(amount float64, events []config.Event) *config.Event {

	var bestEvent *config.Event

	for _, event := range events {
		ev := event
		if ev.Match == "exact" && ev.Amount == amount {
			return &event
		}

		if ev.Match == "minimum" && ev.Amount <= amount {
			if bestEvent == nil || ev.Amount > bestEvent.Amount {
				bestEvent = &ev
			}
		}
	}

	return bestEvent
}
