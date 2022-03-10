package companion

import (
	"fmt"
	"testing"
)

func TestConfig_Matches(t *testing.T) {

	c := Config{Events: []Event{
		{
			Type: EventTypeBits,
			Min:  0,
			Max:  100,
		},
		{
			Type: EventTypeBits,
			Min:  200,
			Max:  250,
		},
		{
			Type: EventTypeDono,
			Min:  20,
			Max:  30,
		},
		{
			Type: EventTypeDono,
			Min:  100,
		},
	}}

	tests := []struct {
		ev   StreamEvent
		want bool
	}{
		{StreamEvent{EventTypeBits, 50}, true},
		{StreamEvent{EventTypeBits, 150}, false},
		{StreamEvent{EventTypeBits, 220}, true},
		{StreamEvent{EventTypeDono, 15}, false},
		{StreamEvent{EventTypeDono, 25}, true},
		{StreamEvent{EventTypeDono, 150}, true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := c.Matches(tt.ev); got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}
