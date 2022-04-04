package companion

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConfig_HasEvent(t *testing.T) {

	sqDefault := SquirtPattern{500 * time.Millisecond}
	sqOne := SquirtPattern{1 * time.Second, 2 * time.Second, 1 * time.Second}
	sqTwo := SquirtPattern{2 * time.Second}
	sqThree := SquirtPattern{3 * time.Second}
	sqFour := SquirtPattern{4 * time.Second}

	c := Config{
		Duration: sqDefault,
		Events: []Event{
			{
				Type:    EventTypeBits,
				Min:     0,
				Max:     100,
				Pattern: sqOne,
			},
			{
				Type:    EventTypeBits,
				Min:     200,
				Max:     250,
				Pattern: sqTwo,
			},
			{
				Type:    EventTypeDono,
				Min:     20,
				Max:     30,
				Pattern: sqThree,
			},
			{
				Type:    EventTypeDono,
				Min:     100,
				Pattern: sqFour,
			},
			{
				Type: EventTypeDono,
				Min:  200,
			},
		}}

	tests := []struct {
		ev          StreamEvent
		want        bool
		wantPattern *SquirtPattern
	}{
		{StreamEvent{EventTypeBits, 50}, true, &sqOne},
		{StreamEvent{EventTypeBits, 150}, false, nil},
		{StreamEvent{EventTypeBits, 220}, true, &sqTwo},
		{StreamEvent{EventTypeDono, 15}, false, nil},
		{StreamEvent{EventTypeDono, 25}, true, &sqThree},
		{StreamEvent{EventTypeDono, 150}, true, &sqFour},
		{StreamEvent{EventTypeDono, 500}, true, &sqDefault},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, p := c.GetEvent(tt.ev)
			if got != tt.want {
				t.Errorf("HasEvent() = %v, want %v", got, tt.want)
			}

			t.Logf("Want: %v", tt.wantPattern)
			t.Logf("Got:  %v", p)

			if !reflect.DeepEqual(p, tt.wantPattern) {
				t.Errorf("HasEvent() = %v, want pattern %v", p, tt.wantPattern)
			}
		})
	}
}

func TestConfig_HasChatTrigger(t *testing.T) {

	sqDefault := SquirtPattern{500 * time.Millisecond}
	sqOne := SquirtPattern{1 * time.Second}
	sqTwo := SquirtPattern{2 * time.Second}
	sqThree := SquirtPattern{3 * time.Second}
	sqFour := SquirtPattern{4 * time.Second}

	c := Config{
		Duration: sqDefault,
		ChatTriggers: []ChatTrigger{
			{
				Role:    ChatRoleMod,
				Message: "modsquirt",
				Pattern: sqOne,
			},
			{
				Role:    ChatRoleSub,
				Message: "subsquirt",
				Pattern: sqTwo,
			},
			{
				Role:    ChatRolePleb,
				Message: "plebsquirt",
				Pattern: sqThree,
			},
			{
				User:    "Steven",
				Message: "stevensquirt",
				Pattern: sqFour,
			},
			{
				User:    "Tom",
				Role:    ChatRoleSub,
				Message: "tomsquirt",
			},
		}}

	tests := []struct {
		message     ChatMessage
		want        bool
		wantPattern *SquirtPattern
	}{
		{ChatMessage{"Steven", ChatRoleMod, "do a modsquirt"}, true, &sqOne},
		{ChatMessage{"Steven", ChatRoleSub, "do a vipsquirt"}, false, nil},
		{ChatMessage{"Steven", ChatRoleSub, "do a subsquirt"}, true, &sqTwo},
		{ChatMessage{"Steven", ChatRolePleb, "do a subsquirt"}, false, nil},
		{ChatMessage{"Steven", ChatRolePleb, "do a plebsquirt"}, true, &sqThree},
		{ChatMessage{"Steven", ChatRolePleb, "do a stevensquirt"}, true, &sqFour},
		{ChatMessage{"Peter", ChatRolePleb, "do a stevensquirt"}, false, nil},
		{ChatMessage{"Tom", ChatRoleMod, "do a tomsquirt"}, true, &sqDefault},
		{ChatMessage{"Tom", ChatRoleSub, "do a tomsquirt"}, true, &sqDefault},
		{ChatMessage{"Tom", ChatRolePleb, "do a tomsquirt"}, false, nil},
		{ChatMessage{"Karl", ChatRoleSub, "do a tomsquirt"}, false, nil},
		{ChatMessage{"Karl", ChatRoleMod, "do a tomsquirt"}, false, nil},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, p := c.GetChatTrigger(tt.message)
			if got != tt.want {
				t.Errorf("HasChatTrigger() = %v, want %v\n%v", got, tt.want, tt.message)
			}

			t.Logf("Want: %v", tt.wantPattern)
			t.Logf("Got:  %v", p)

			if !reflect.DeepEqual(p, tt.wantPattern) {
				t.Errorf("HasChatTrigger() = %v, want pattern %v\n%v", p, tt.wantPattern, tt.message)
			}
		})
	}
}
