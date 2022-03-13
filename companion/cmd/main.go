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
	time.Sleep(100 * time.Millisecond)
	log.Println("Squirtianna companion starting...")

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
	log.Println("Config loaded:")
	conf.Dump(os.Stdout)

	events := make(chan companion.StreamEvent)
	messages := make(chan companion.ChatMessage)

	if conf.Streamlabs != "" {
		sl := streamlabs.New(conf.Streamlabs)
		sl.Connect(events, messages)
	}

	if conf.Twitch != "" {
		//tc := twitchchat.New(conf.Twitch, twitchchat.WithFdgt(), twitchchat.WithFaker(3270*time.Millisecond))
		tc := twitchchat.New(conf.Twitch)
		tc.Connect(events, messages)
	}

	var squirters squirter.Squirters

	if len(conf.Squirters) > 0 {
		for _, s := range conf.Squirters {
			squirters = append(squirters, squirter.New(s))
		}
	} else {
		squirters = squirter.Find()
	}

	for _, s := range squirters {
		log.Printf("Use %v", s)
	}

	timeout := time.Now()

	for {
		select {
		case m := <-messages:
			if conf.HasChatTrigger(m) && time.Now().After(timeout) {
				log.Printf("Message from %v: %v -> Squirt for %v", m.User, m.Message, conf.Duration)
				squirters.Squirt(conf.Duration)
				timeout = time.Now().Add(conf.Cooldown + conf.Duration)
			}
		case e := <-events:
			if conf.HasEvent(e) && time.Now().After(timeout) {
				log.Printf("%v of %v: Squirt for %v", e.EventType, e.Amount, conf.Duration)
				squirters.Squirt(conf.Duration)
				timeout = time.Now().Add(conf.Cooldown + conf.Duration)
			}
		}
	}

}
