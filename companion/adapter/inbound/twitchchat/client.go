package twitchchat

import (
	"companion"
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const Fdgt = "irc.fdgt.dev:6697"

type Twitch interface {
}

type Opt func(chat *twitchChat)

func New(channel string, opts ...Opt) companion.StreamEventSource {
	t := &twitchChat{
		channel: channel,
		client:  twitch.NewAnonymousClient(),
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func WithFdgt() Opt {
	return WithIrcServer(Fdgt)
}
func WithIrcServer(server string) Opt {
	return func(chat *twitchChat) {
		chat.client.IrcAddress = server
	}
}
func WithFaker(interval time.Duration) Opt {
	return func(chat *twitchChat) {
		go chat.FakerShaker(interval)
	}
}
func (t *twitchChat) FakerShaker(interval time.Duration) {

	if t.client.IrcAddress != Fdgt {
		log.Fatalf("Can only use fakes on %v!", Fdgt)
	}

	for {
		time.Sleep(interval)
		switch rand.Intn(5) {
		case 0:
			t.client.Say("channel", fmt.Sprintf("bits --bitscount %v Woohoo!", rand.Intn(500)))
			//fallthrough
		case 1:
			t.client.Say("channel", fmt.Sprintf("subgift --tier %v --username glEnd2", rand.Intn(3)+1))
			//fallthrough
		case 2:
			t.client.Say("channel", fmt.Sprintf("submysterygift --count %v --username zebiniasis", rand.Intn(15)+1))
			//fallthrough
		case 3:
			t.client.Say("channel", fmt.Sprintf("subscription --tier %v", rand.Intn(3)+1))
			//fallthrough
		case 4:
			t.client.Say("channel", fmt.Sprintf("resubscription --tier %v", rand.Intn(3)+1))
		}

	}

}

type twitchChat struct {
	client  *twitch.Client
	channel string
}

func (t *twitchChat) Connect(events chan<- companion.StreamEvent, messages chan<- companion.ChatMessage) error {
	t.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.Bits > 0 {
			events <- companion.StreamEvent{
				EventType: companion.EventTypeBits,
				Amount:    message.Bits,
			}
		} else {
			role := companion.ChatRolePleb
			if _, exists := message.User.Badges["subscriber"]; exists {
				role = companion.ChatRoleSub
			}
			if _, exists := message.User.Badges["moderator"]; exists {
				role = companion.ChatRoleMod
			}
			if _, exists := message.User.Badges["broadcaster"]; exists {
				role = companion.ChatRoleMod
			}
			messages <- companion.ChatMessage{
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
			events <- companion.StreamEvent{
				EventType: companion.EventTypeSub,
				Amount:    count,
			}

		case "sub":
			fallthrough

		case "resub":
			plan, _ := strconv.Atoi(message.MsgParams["msg-param-sub-plan"])
			months, _ := strconv.Atoi(message.MsgParams["msg-param-cumulative-months"])
			events <- companion.StreamEvent{
				EventType: planToEventType(plan),
				Amount:    months,
			}

		case "subgift":
			events <- companion.StreamEvent{
				EventType: companion.EventTypeSub,
				Amount:    1,
			}
		}
	})

	t.client.Join(t.channel)

	go func() {
		log.Println("Twitch connected")
		_ = t.client.Connect()
	}()

	return nil
}

func planToEventType(plan int) companion.EventType {
	switch plan {
	case 3000:
		return companion.EventTypeT3Sub
	case 2000:
		return companion.EventTypeT2Sub
	default:
		return companion.EventTypeT1Sub
	}
}
