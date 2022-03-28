package streamlabs

import (
	"companion"
	"fmt"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"log"
	"strconv"
	"time"
)

func New(token, currency string) companion.StreamEventSource {
	return &streamlabs{
		token:    token,
		currency: currency,
	}
}

type streamlabs struct {
	token    string
	currency string
}

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (s *streamlabs) Connect(events chan<- companion.StreamEvent, messages chan<- companion.ChatMessage) error {
	websocketTransport := transport.GetDefaultWebsocketTransport()
	websocketTransport.PingInterval = 20 * time.Second

	client, err := gosocketio.Dial(
		gosocketio.GetUrl("sockets.streamlabs.com", 443, true)+"&token="+s.token,
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe to create client: %w", err)
	}

	err = client.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Printf("Streamlabs connected")
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to connects: %w", err)
	}

	err = client.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Printf("Streamlabs disconnected")
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to disconnects: %w", err)
	}

	err = client.On("event", func(c *gosocketio.Channel, data Ev) {
		if data.Type == "donation" && data.Message[0].Currency == s.currency {
			amount := parseAmount(data.Message[0].Amount) * 100
			log.Println("Amount:", amount)
			events <- companion.StreamEvent{
				EventType: companion.EventTypeDono,
				Amount:    amount,
			}
		}
	})

	if err != nil {
		return fmt.Errorf("failed to subscribe to events: %w", err)
	}

	return nil
}

func parseAmount(input interface{}) int {
	asString := fmt.Sprint(input)
	num, err := strconv.Atoi(asString)
	if err != nil {
		return -1
	}
	return num
}

type Ev struct {
	For     string `json:"for"`
	Type    string `json:"type"`
	Message []struct {
		Amount   interface{} `json:"amount"`
		Currency string      `json:"currency"`
	} `json:"message"`
}
