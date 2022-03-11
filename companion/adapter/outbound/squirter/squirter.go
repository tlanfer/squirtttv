package squirter

import (
	"fmt"
	"github.com/hashicorp/mdns"
	"net/http"
	"time"
)

type Squirters []Squirter

func (s *Squirters) Squirt(duration time.Duration) {
	for _, sq := range *s {
		go sq.Squirt(duration)
	}
}

type Squirter interface {
	Squirt(duration time.Duration)
}

func Find() Squirters {

	var squirters []Squirter

	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			h := entry.AddrV4.String()
			squirters = append(squirters, &squirter{
				host: h,
			})
		}
	}()
	params := mdns.DefaultParams("_squirtianna._tcp")
	params.Entries = entriesCh
	params.DisableIPv6 = true
	_ = mdns.Query(params)
	close(entriesCh)

	return squirters
}

func New(host string) Squirter {
	return &squirter{host: host}
}

type squirter struct {
	host string
}

func (s *squirter) String() string {
	return fmt.Sprintf("squirter [%v]", s.host)
}

func (s *squirter) Squirt(duration time.Duration) {
	http.DefaultClient.Timeout = 1 * time.Second
	http.Get(fmt.Sprintf("http://%v/squirt?duration=%v", s.host, duration.Milliseconds()))
}
