package squirter

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Squirter struct {
	Host string
}

func (s Squirter) String() string {
	return fmt.Sprintf("squirter [%v]", s.Host)
}

func (s Squirter) Squirt(duration time.Duration) {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DisableKeepAlives = true
	c := &http.Client{
		Transport: t,
		Timeout:   5 * time.Second,
	}
	_, err := c.Get(fmt.Sprintf("http://%v/squirt?duration=%v", s.Host, duration.Milliseconds()))
	if err != nil {
		log.Printf("Failed to send event to %v: %v", s, err)
	}
}

func (s Squirter) IsSquirter() bool {

	resp, err := http.Get(fmt.Sprintf("http://%v/identify", s.Host))
	if err != nil {
		log.Printf("failed to identify squirter on %v: %v", s.Host, err)
		return false
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.Header.Get("Server") != "squirtttv/2.0" {
		log.Printf("device at %v is not a squirter", s.Host)
		return false
	}

	return true
}
