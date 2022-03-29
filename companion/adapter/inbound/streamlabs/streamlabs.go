package streamlabs

import (
	"companion"
	"fmt"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"log"
	"strconv"
	"strings"
	"time"
)

func New(token, currency string, converter companion.CurrencyConverter) companion.StreamEventSource {
	return &streamlabs{
		token:     token,
		currency:  currency,
		converter: converter,
	}
}

type streamlabs struct {
	token     string
	currency  string
	converter companion.CurrencyConverter
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
		if data.Type == "donation" {
			amount := parseAmount(data.Message[0].Amount)
			sourceCurrency := strings.ToLower(data.Message[0].Currency)
			if sourceCurrency != s.currency {
				converted := s.converter.Convert(int(amount), sourceCurrency, s.currency)
				log.Printf("Streamlabs donation: %.2f %v converted to %.2f %v", float32(amount)/100, sourceCurrency, float32(converted)/100, s.currency)
				amount = converted
			} else {
				log.Printf("Streamlabs donation: %.2f %v", float32(amount)/100, s.currency)
			}
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
	num, err := strconv.ParseFloat(asString, 32)
	if err != nil {
		return -1
	}
	return int(num * 100)
}

type Ev struct {
	For     string `json:"for"`
	Type    string `json:"type"`
	Message []struct {
		Amount   interface{} `json:"amount"`
		Currency string      `json:"currency"`
	} `json:"message"`
}
