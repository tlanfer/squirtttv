package main

import (
	"companion"
	"companion/adapter/inbound/streamlabs"
	"companion/adapter/inbound/twitchchat"
	"companion/adapter/inbound/yamlconfig"
	"companion/adapter/outbound/squirter"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	filename = "config.yaml"
	example  = "config.example.yaml"
)

func main() {

	loader := yamlconfig.New(filename, example)
	conf, err := loader.Load()
	if err != nil {
		if err == companion.ErrConfigNotFound {
			fmt.Printf("No %v found, please create one. You can use %v as an example.", filename, example)
			os.Exit(1)
		} else {
			fmt.Println("Failed to load config:", err.Error())
			os.Exit(1)
		}
	}

	events := make(chan companion.StreamEvent)

	if conf.Streamlabs != "" {
		sl := streamlabs.New(conf.Streamlabs)
		sl.Connect(events)
	}

	if conf.Twitch != "" {
		//tc := twitchchat.New(conf.Twitch, twitchchat.WithFdgt(), twitchchat.WithFaker(3270*time.Millisecond))
		tc := twitchchat.New(conf.Twitch)
		tc.Connect(events)
	}

	squirters := squirter.Find()

	timeout := time.Now()

	for event := range events {
		if time.Now().After(timeout) && conf.Matches(event) {
			timeout = time.Now().Add(conf.Cooldown)
			log.Printf("Got some %v (%v) -> Squirt for %v", event.EventType, event.Amount, conf.Duration)
			squirters.Squirt(conf.Duration)
		} else {
			log.Printf("Ignore %v (%v)", event.EventType, event.Amount)
		}
	}

}
