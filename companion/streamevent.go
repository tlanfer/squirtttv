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
	return fmt.Sprintf("[ %v | %v ]", s.EventType, s.Amount)
}

type StreamEventSource interface {
	Connect(events chan<- StreamEvent, messages chan<- ChatMessage) error
}
