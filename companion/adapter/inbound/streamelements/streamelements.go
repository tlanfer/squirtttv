package streamelements

import (
	"companion"
	"encoding/json"
	"fmt"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"log"
	"strings"
	"time"
)

func New(token, currency string, converter companion.CurrencyConverter) companion.StreamEventSource {
	return &streamelements{
		token:     token,
		currency:  currency,
		converter: converter,
	}
}

type streamelements struct {
	token     string
	currency  string
	converter companion.CurrencyConverter
}

func (s streamelements) Connect(events chan<- companion.StreamEvent, messages chan<- companion.ChatMessage) error {

	websocketTransport := transport.GetDefaultWebsocketTransport()
	websocketTransport.PingInterval = 5 * time.Second

	client, err := gosocketio.Dial(
		gosocketio.GetUrl("realtime.streamelements.com", 443, true),
		websocketTransport,
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe to create client: %w", err)
	}

	_ = client.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		err := c.Emit("authenticate", map[string]string{
			"method": "jwt",
			"token":  s.token,
		})
		if err != nil {
			log.Println("failed to authenticate with streamelements: ", err)
		}
	})

	client.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Printf("StreamElements disconnected")
	})

	client.On("authenticated", func(c *gosocketio.Channel) {
		log.Println("StreamElements connected")
	})

	client.On("unauthorized", func(c *gosocketio.Channel) {
		log.Println("Unauthorized")
	})

	if err := client.On("event", func(c *gosocketio.Channel, data string) {
		event := Event{}
		if err := unmarshalEvent(data, &event); err != nil {
			log.Println("failed to unmarshal event: ", err)
			return
		}

		if event.Type != "tip" {
			return
		}

		fromCurrency := strings.ToLower(event.Data.Currency)
		toCurrency := strings.ToLower(s.currency)
		originalAmount := int(event.Data.Amount * 100)
		finalAmount := originalAmount

		if fromCurrency == toCurrency {
			log.Printf("StreamElements donation: %.2f %v", float32(finalAmount)/100, fromCurrency)
		} else {
			finalAmount = s.converter.Convert(originalAmount, event.Data.Currency, s.currency)
			log.Printf("StreamElements donation: %.2f %v converted to %.2f %v", float32(originalAmount)/100, fromCurrency, float32(finalAmount)/100, toCurrency)
		}

		events <- companion.StreamEvent{
			EventType: companion.EventTypeDono,
			Amount:    finalAmount,
		}

	}); err != nil {
		return fmt.Errorf("failed to subscribe to event: %w", err)
	}

	return nil
}

func unmarshalEvent(data string, dto any) error {
	var parsed []json.RawMessage
	if err := json.Unmarshal([]byte(fmt.Sprintf("[%s]", data)), &parsed); err != nil {
		return err
	}
	return json.Unmarshal(parsed[0], &dto)
}

type Event struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Data    struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"data"`
}
