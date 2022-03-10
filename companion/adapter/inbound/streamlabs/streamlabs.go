package streamlabs

import (
	"companion"
	"fmt"
	gosocketio "github.com/mtfelian/golang-socketio"
	"github.com/mtfelian/golang-socketio/transport"
	"log"
	"time"
)

func New(token string) companion.StreamEventSource {
	return &streamlabs{
		token: token,
	}
}

type streamlabs struct {
	token string
}

func (s *streamlabs) Connect(events chan<- companion.StreamEvent) error {
	websocketTransport := transport.DefaultWebsocketTransport()
	websocketTransport.PingInterval = 20 * time.Second

	client, err := gosocketio.Dial(
		gosocketio.AddrWebsocket("sockets.streamlabs.com", 443, true)+"&token="+s.token,
		websocketTransport,
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe to create client: %w", err)
	}

	err = client.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Printf("StreamlabsConfig connected")
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to connects: %w", err)
	}

	err = client.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Printf("StreamlabsConfig disconnected")
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to disconnects: %w", err)
	}

	err = client.On("event", func(c *gosocketio.Channel, data Ev) {
		if data.Type == "donation" {
			amount := int(data.Message[0].Amount * 100)
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

type Ev struct {
	For     string `json:"for"`
	Type    string `json:"type"`
	Message []struct {
		Amount   float32 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"message"`
}
