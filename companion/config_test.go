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
	sqFive := SquirtPattern{1234 * time.Millisecond}
	sqSix := SquirtPattern{2345 * time.Millisecond}
	sqSeven := SquirtPattern{3456 * time.Millisecond}
	sq8 := SquirtPattern{8 * time.Millisecond}

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
			{
				Type:    EventTypeT3Sub,
				Pattern: sqFive,
			},
			{
				Type:    EventTypeT3Sub,
				Pattern: sqSix,
				Min:     15,
			},
			{
				Type:    EventTypeT1Sub,
				Pattern: sqSeven,
				Min:     10,
				Max:     10,
			},
			{
				Type:    EventTypeT2Sub,
				Pattern: sq8,
				Min:     2,
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
		{StreamEvent{EventTypeT3Sub, 10}, true, &sqFive},
		{StreamEvent{EventTypeT3Sub, 20}, true, &sqSix},
		{StreamEvent{EventTypeT2Sub, 1}, false, nil},
		{StreamEvent{EventTypeT1Sub, 10}, true, &sqSeven},
		{StreamEvent{EventTypeT1Sub, 11}, false, nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.ev.EventType, tt.ev.Amount), func(t *testing.T) {
			got, p := c.GetEvent(tt.ev)
			if got != tt.want {
				t.Errorf("HasEvent() = %v, want %v", got, tt.want)
			}

			t.Logf("Event: %v", tt.ev)
			t.Logf("Want: %v", tt.wantPattern)
			t.Logf("Got:  %v", p)

			if !reflect.DeepEqual(p, tt.wantPattern) {
				t.Errorf("HasEvent() = %v, want pattern %v", p, tt.wantPattern)
			}
		})
	}
}

func TestConfig_AlasConfig(t *testing.T) {

	def := SquirtPattern{1 * time.Second}
	twoSeconds := SquirtPattern{2 * time.Second}
	threeSeconds := SquirtPattern{3 * time.Second}
	fourSeconds := SquirtPattern{4 * time.Second}
	onOff5times := SquirtPattern{1 * time.Second, 200 * time.Millisecond, 1 * time.Second, 200 * time.Millisecond, 1 * time.Second, 200 * time.Millisecond, 1 * time.Second, 200 * time.Millisecond, 1 * time.Second}
	t2 := SquirtPattern{500 * time.Millisecond, 200 * time.Millisecond, 500 * time.Millisecond}
	t3 := SquirtPattern{500 * time.Millisecond, 200 * time.Millisecond, 500 * time.Millisecond, 200 * time.Millisecond, 500 * time.Millisecond}

	c := Config{
		Duration: def,
		Events: []Event{
			{
				Type: EventTypeBits,
				Min:  500,
			}, {
				Type:    EventTypeBits,
				Min:     1000,
				Pattern: twoSeconds,
			}, {
				Type:    EventTypeBits,
				Min:     2000,
				Pattern: threeSeconds,
			}, {
				Type:    EventTypeBits,
				Min:     1500,
				Max:     1500,
				Pattern: fourSeconds,
			}, {
				Type:    EventTypeBits,
				Min:     5000,
				Pattern: onOff5times,
			}, {
				Type: EventTypeDono,
				Min:  500,
			}, {
				Type:    EventTypeDono,
				Min:     1000,
				Pattern: twoSeconds,
			}, {
				Type:    EventTypeDono,
				Min:     2000,
				Pattern: threeSeconds,
			}, {
				Type:    EventTypeDono,
				Min:     5000,
				Pattern: onOff5times,
			}, {
				Type:    EventTypeT2Sub,
				Pattern: t2,
			}, {
				Type:    EventTypeT3Sub,
				Pattern: t3,
			},
		}}

	tests := []struct {
		ev          StreamEvent
		want        bool
		wantPattern *SquirtPattern
	}{
		{StreamEvent{EventTypeBits, 1500}, true, &fourSeconds},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.ev.EventType, tt.ev.Amount), func(t *testing.T) {
			got, p := c.GetEvent(tt.ev)
			if got != tt.want {
				t.Errorf("HasEvent() = %v, want %v", got, tt.want)
			}

			t.Logf("Event: %v", tt.ev)
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
