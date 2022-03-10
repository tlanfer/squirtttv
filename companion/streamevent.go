package companion

import "fmt"

type EventType string

const (
	EventTypeBits = "bits"
	EventTypeDono = "dono"
	EventTypeSub  = "subs"
)

type StreamEvent struct {
	EventType EventType
	Amount    int
}

func (s StreamEvent) String() string {
	t := ""
	switch s.EventType {
	case EventTypeDono:
		return fmt.Sprintf("[ Dono: %.2f usd ]", float64(s.Amount)/100)
	case EventTypeBits:
		return fmt.Sprintf("[ Bits: %d bits ]", s.Amount)
	case EventTypeSub:
		return fmt.Sprintf("[ Gift: %d subs ]", s.Amount)
	}
	return t
}

type StreamEventSource interface {
	Connect(events chan<- StreamEvent) error
}
