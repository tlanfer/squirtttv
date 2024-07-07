package streamlabs

import (
	"companion/internal"
	"companion/internal/config"
	"companion/internal/state"
	"fmt"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"log"
	"strconv"
	"time"
)

func New(events chan<- internal.StreamEvent) {
	c := config.Get()
	s := streamlabs{
		token:  c.Settings.Streamlabs,
		events: events,
	}

	config.Subscribe(func(c config.Config) {
		if c.Settings.Streamlabs == s.token {
			return
		}
		log.Println("Streamlabs token changed")
		s.token = c.Settings.Streamlabs
		if s.client != nil && s.client.IsAlive() {
			s.client.Close()
			time.Sleep(1 * time.Second)
		}
		go s.Connect()
	})
	go s.Connect()
}

type streamlabs struct {
	token    string
	client   *gosocketio.Client
	events   chan<- internal.StreamEvent
	messages chan<- internal.ChatMessage
}

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (s *streamlabs) Connect() {
	if s.token == "" {
		return
	}

	websocketTransport := transport.GetDefaultWebsocketTransport()
	websocketTransport.PingInterval = 5 * time.Second

	client, err := gosocketio.Dial(gosocketio.GetUrl("sockets.streamlabs.com", 443, true)+"&token="+s.token, websocketTransport)
	s.client = client

	if err != nil {
		log.Printf("failed to subscribe to create client: %v", err)
		return
	}

	err = client.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Printf("Streamlabs connected")
		state.StreamlabsConnected = true
	})
	if err != nil {
		log.Printf("failed to subscribe to connects: %v", err)
		return
	}

	err = client.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Printf("Streamlabs disconnected")
		state.StreamlabsConnected = false
		time.Sleep(10 * time.Second)
		go s.Connect()
	})
	if err != nil {
		log.Printf("failed to subscribe to disconnects: %v", err)
		return
	}

	err = client.On("event", func(c *gosocketio.Channel, data Ev) {
		if data.Type == "donation" {
			amount := parseAmount(data.Message[0].Amount)
			//sourceCurrency := strings.ToLower(data.Message[0].Currency)
			//if sourceCurrency != s.currency {
			//	converted := s.converter.Convert(int(amount), sourceCurrency, s.currency)
			//	log.Printf("Streamlabs donation: %.2f %v converted to %.2f %v", float32(amount)/100, sourceCurrency, float32(converted)/100, s.currency)
			//	amount = converted
			//} else {
			//	log.Printf("Streamlabs donation: %.2f %v", float32(amount)/100, s.currency)
			//}
			s.events <- internal.StreamEvent{
				EventType: internal.EventTypeDono,
				Amount:    amount,
			}
		}
	})

	if err != nil {
		log.Printf("failed to subscribe to events: %v", err)
		return
	}
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
