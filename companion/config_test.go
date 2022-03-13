package companion

import (
	"fmt"
	"testing"
)

func TestConfig_HasEvent(t *testing.T) {

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
			if got := c.HasEvent(tt.ev); got != tt.want {
				t.Errorf("HasEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_HasChatTrigger(t *testing.T) {
	c := Config{ChatTriggers: []ChatTrigger{
		{
			Role:    ChatRoleMod,
			Message: "modsquirt",
		},
		{
			Role:    ChatRoleSub,
			Message: "subsquirt",
		},
		{
			Role:    ChatRolePleb,
			Message: "plebsquirt",
		},
		{
			User:    "Steven",
			Message: "stevensquirt",
		},
		{
			User:    "Tom",
			Role:    ChatRoleSub,
			Message: "tomsquirt",
		},
	}}

	tests := []struct {
		message ChatMessage
		want    bool
	}{
		{ChatMessage{"Steven", ChatRoleMod, "do a modsquirt"}, true},
		{ChatMessage{"Steven", ChatRoleSub, "do a vipsquirt"}, false},
		{ChatMessage{"Steven", ChatRoleSub, "do a subsquirt"}, true},
		{ChatMessage{"Steven", ChatRolePleb, "do a subsquirt"}, false},
		{ChatMessage{"Steven", ChatRolePleb, "do a plebsquirt"}, true},
		{ChatMessage{"Steven", ChatRolePleb, "do a stevensquirt"}, true},
		{ChatMessage{"Peter", ChatRolePleb, "do a stevensquirt"}, false},
		{ChatMessage{"Tom", ChatRoleMod, "do a tomsquirt"}, true},
		{ChatMessage{"Tom", ChatRoleSub, "do a tomsquirt"}, true},
		{ChatMessage{"Tom", ChatRolePleb, "do a tomsquirt"}, false},
		{ChatMessage{"Karl", ChatRoleSub, "do a tomsquirt"}, false},
		{ChatMessage{"Karl", ChatRoleMod, "do a tomsquirt"}, false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := c.HasChatTrigger(tt.message); got != tt.want {
				t.Errorf("HasChatTrigger() = %v, want %v\n%v", got, tt.want, tt.message)
			}
		})
	}
}
