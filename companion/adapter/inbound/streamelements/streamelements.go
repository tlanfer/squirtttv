package streamelements

import (
	"companion"
	"encoding/json"
	"fmt"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"log"
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

	client.On("authenticated", func(c *gosocketio.Channel, data Authenticated) {
		log.Println("Authenticated: ", data.Message)
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

		eventAmount := int(event.Data.Amount * 100)
		convertedAmount := s.converter.Convert(eventAmount, event.Data.Currency, s.currency)

		events <- companion.StreamEvent{
			EventType: companion.EventTypeDono,
			Amount:    convertedAmount,
		}

	}); err != nil {
		return fmt.Errorf("failed to subscribe to event: %w", err)
	}

	if err := client.On("event:test", func(c *gosocketio.Channel, data string) {
		log.Println("event:test: ", data)
	}); err != nil {
		return fmt.Errorf("failed to subscribe to event:test: %w", err)
	}

	//if err := client.On("event:update", func(c *gosocketio.Channel, data string) {
	//	log.Println("event:update: ", data)
	//}); err != nil {
	//	return fmt.Errorf("failed to subscribe to event:update: %w", err)
	//}

	return nil
}

func unmarshalEvent(data string, dto any) error {
	var parsed []json.RawMessage
	if err := json.Unmarshal([]byte(fmt.Sprintf("[%s]", data)), &parsed); err != nil {
		return err
	}
	return json.Unmarshal(parsed[0], &dto)
}

type Authenticated struct {
	Channel  string `json:"channelId"`
	ClientId string `json:"clientId"`
	Message  string `json:"message"`
	Project  string `json:"project"`
}

type Event struct {
	Type      string    `json:"type"`
	Provider  string    `json:"provider"`
	Channel   string    `json:"channel"`
	CreatedAt time.Time `json:"createdAt"`
	Data      struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
		Username string  `json:"username"`
		TipID    string  `json:"tipId"`
		Message  string  `json:"message"`
		Avatar   string  `json:"avatar"`
	} `json:"data"`
	ID                 string    `json:"_id"`
	UpdatedAt          time.Time `json:"updatedAt"`
	ActivityID         string    `json:"activityId"`
	SessionEventsCount int       `json:"sessionEventsCount"`
}
