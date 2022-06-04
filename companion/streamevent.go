package companion

import "fmt"

type EventType string

const (
	EventTypeBits  = "bits"
	EventTypeDono  = "dono"
	EventTypeSub   = "subs"
	EventTypeT1Sub = "tier1"
	EventTypeT2Sub = "tier2"
	EventTypeT3Sub = "tier3"
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
