package twitchchat

import (
	"companion/internal"
	"companion/internal/config"
	"companion/internal/state"
	"github.com/gempir/go-twitch-irc/v3"
	"log"
	"strconv"
	"time"
)

const Fdgt = "irc.fdgt.dev:6697"

type Twitch interface {
}

type Opt func(chat *twitchChat)

func New(events chan<- internal.StreamEvent, messages chan<- internal.ChatMessage, opts ...Opt) {
	c := config.Get()
	t := &twitchChat{
		channel:  c.Settings.Twitch,
		client:   twitch.NewAnonymousClient(),
		events:   events,
		messages: messages,
	}
	for _, opt := range opts {
		opt(t)
	}

	t.RegisterHandlers()

	config.Subscribe(func(c config.Config) {
		if c.Settings.Twitch == t.channel {
			return
		}
		t.channel = c.Settings.Twitch
		_ = t.client.Disconnect()
		time.Sleep(1 * time.Second)
		go t.Connect()
	})

	go t.Connect()
}

type twitchChat struct {
	client   *twitch.Client
	channel  string
	events   chan<- internal.StreamEvent
	messages chan<- internal.ChatMessage
}

func (t *twitchChat) RegisterHandlers() {
	t.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.Bits > 0 {
			t.events <- internal.StreamEvent{
				EventType: internal.EventTypeBits,
				Amount:    message.Bits,
			}
		} else {
			role := internal.ChatRolePleb
			if _, exists := message.User.Badges["subscriber"]; exists {
				role = internal.ChatRoleSub
			}
			if _, exists := message.User.Badges["moderator"]; exists {
				role = internal.ChatRoleMod
			}
			if _, exists := message.User.Badges["broadcaster"]; exists {
				role = internal.ChatRoleMod
			}
			t.messages <- internal.ChatMessage{
				User:    message.User.Name,
				Role:    role,
				Message: message.Message,
			}
		}
	})

	t.client.OnUserNoticeMessage(func(message twitch.UserNoticeMessage) {
		switch message.MsgID {
		case "submysterygift":
			count, _ := strconv.Atoi(message.MsgParams["msg-param-mass-gift-count"])
			t.events <- internal.StreamEvent{
				EventType: internal.EventTypeGift,
				Amount:    count,
			}

		case "sub":
			fallthrough

		case "resub":
			plan, _ := strconv.Atoi(message.MsgParams["msg-param-sub-plan"])
			months, _ := strconv.Atoi(message.MsgParams["msg-param-cumulative-months"])
			t.events <- internal.StreamEvent{
				EventType: planToEventType(plan),
				Amount:    months,
			}

		case "subgift":
			t.events <- internal.StreamEvent{
				EventType: internal.EventTypeGift,
				Amount:    1,
			}
		}
	})
}

func (t *twitchChat) Connect() {
	if t.channel == "" {
		log.Println("No twitch channel set")
		return
	}
	t.client.Join(t.channel)
	log.Println("Twitch connected")
	state.TwitchConnected = true
	_ = t.client.Connect()
	state.TwitchConnected = false
}

func planToEventType(plan int) internal.EventType {
	switch plan {
	case 3000:
		return internal.EventTypeT3Sub
	case 2000:
		return internal.EventTypeT2Sub
	default:
		return internal.EventTypeT1Sub
	}
}
