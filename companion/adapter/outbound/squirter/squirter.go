package squirter

import (
	"context"
	"fmt"
	"github.com/grandcat/zeroconf"
	"log"
	"net/http"
	"time"
)

type Squirters struct {
	all          []Squirter
	allowInvalid bool
}

func NewSquirters(allowInvalid bool) Squirters {
	return Squirters{allowInvalid: !allowInvalid}
}

func (s *Squirters) Squirt(duration time.Duration) {
	for _, sq := range s.all {
		go sq.Squirt(duration)
	}
}

type Squirter interface {
	Squirt(duration time.Duration)
}

func (s *Squirters) Add(host string) {
	sq := squirter{host: host}
	if sq.isAvailable() || s.allowInvalid {
		s.all = append(s.all, &sq)
		log.Printf("Squirter at %v added", host)
	} else {
		log.Printf("Squirter at %v invalid", host)
	}
}

func (s *Squirters) Find() {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Println("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			if len(entry.AddrIPv4) < 1 {
				continue
			}
			h := entry.AddrIPv4[0].String()

			s.Add(h)
		}
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = resolver.Browse(ctx, "_squirtttv._tcp", "", entries)
	if err != nil {
		log.Println("Failed to browse:", err.Error())
	}

	<-ctx.Done()
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
	http.DefaultClient.Timeout = 5 * time.Second
	_, err := http.Get(fmt.Sprintf("http://%v/squirt?duration=%v", s.host, duration.Milliseconds()))
	if err != nil {
		log.Printf("Failed to send event to %v: %v", s, err)
	}
}

func (s *squirter) isAvailable() bool {

	resp, err := http.Get(fmt.Sprintf("http://%v/identify", s.host))
	if err != nil {
		log.Printf("failed to identify squirter on %v: %v", s.host, err)
		return false
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.Header.Get("Server") != "squirtttv/2.0" {
		log.Printf("device at %v is not a squirter", s.host)
		return false
	}

	return true
}
