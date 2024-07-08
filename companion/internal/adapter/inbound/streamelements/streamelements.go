package streamelements

import (
	"companion/internal"
	"companion/internal/adapter/outbound/exchangerate"
	"companion/internal/config"
	"companion/internal/state"
	"encoding/json"
	"fmt"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"log"
	"time"
)

func New(events chan<- internal.StreamEvent) {
	c := config.Get()
	s := streamelements{
		token:  c.Settings.Streamelements,
		events: events,
	}
	config.Subscribe(func(c config.Config) {
		if c.Settings.Streamelements == s.token {
			return
		}
		s.token = c.Settings.Streamelements
		if s.client != nil && s.client.IsAlive() {
			s.client.Close()
			time.Sleep(1 * time.Second)
		}
		go s.Connect()
	})
	go s.Connect()
}

type streamelements struct {
	token    string
	currency string
	events   chan<- internal.StreamEvent
	client   *gosocketio.Client
}

func (s streamelements) Connect() {
	if s.token == "" {
		log.Println("StreamElements token not set, skipping connection")
		return
	}

	websocketTransport := transport.GetDefaultWebsocketTransport()
	websocketTransport.PingInterval = 5 * time.Second

	client, err := gosocketio.Dial(
		gosocketio.GetUrl("realtime.streamelements.com", 443, true),
		websocketTransport,
	)
	s.client = client

	if err != nil {
		log.Println("failed to connect to streamelements: ", err)
		return
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

	_ = client.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Printf("StreamElements disconnected")
		state.StreamElementsConnected = false
	})

	_ = client.On("authenticated", func(c *gosocketio.Channel) {
		log.Println("StreamElements connected")
		state.StreamElementsConnected = true
	})

	_ = client.On("unauthorized", func(c *gosocketio.Channel) {
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

		log.Println("Donation received: ", event.Data.Amount)
		originalAmount := event.Data.Amount
		finalAmount := exchangerate.Convert(originalAmount, "usd")
		finalAmount = float64(int(finalAmount*100)) / 100
		log.Println("Converted donation amount: ", finalAmount)

		s.events <- internal.StreamEvent{
			EventType: internal.EventTypeDono,
			Amount:    finalAmount,
		}

	}); err != nil {
		log.Println("failed to subscribe to event: ", err)
		return
	}
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
