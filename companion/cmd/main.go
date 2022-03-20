package main

import (
	"companion"
	"companion/adapter/inbound/streamlabs"
	"companion/adapter/inbound/twitchchat"
	"companion/adapter/inbound/yamlconfig"
	"companion/adapter/outbound/squirter"
	"companion/adapter/outbound/trayicon"
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

	ui := trayicon.New()

	loader := yamlconfig.New(filename, example)
	conf, err := loader.Load()
	if err != nil {
		if err == companion.ErrConfigNotFound {
			ui.ErrorMessage("No %v found, please create one. You can use %v as an example.", filename, example)
			//fmt.Printf("No %v found, please create one. You can use %v as an example.", filename, example)
			os.Exit(1)
		} else {
			ui.ErrorMessage("Failed to load config: %s", err.Error())
			//fmt.Println("Failed to load config:", err.Error())
			os.Exit(1)
		}
	}

	events := make(chan companion.StreamEvent)
	messages := make(chan companion.ChatMessage)

	if conf.Streamlabs != "" {
		sl := streamlabs.New(conf.Streamlabs)
		err := sl.Connect(events, messages)

		if err != nil {
			ui.ErrorMessage("Failed to connect to streamlabs: %s", err.Error())
		} else {
			ui.SetStreamlabsConnected(true)
		}
	}

	if conf.Twitch != "" {
		//tc := twitchchat.New(conf.Twitch, twitchchat.WithFdgt(), twitchchat.WithFaker(3270*time.Millisecond))
		tc := twitchchat.New(conf.Twitch)
		err := tc.Connect(events, messages)

		if err != nil {
			ui.ErrorMessage("Failed to connect to twitch: %s", err.Error())
		} else {
			ui.SetTwitchConnected(true)
		}
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
				ui.SetActive(conf.Duration)
				squirters.Squirt(conf.Duration)
				timeout = time.Now().Add(conf.Cooldown + conf.Duration)
			}
		case e := <-events:
			if conf.HasEvent(e) && time.Now().After(timeout) {
				log.Printf("%v of %v: Squirt for %v", e.EventType, e.Amount, conf.Duration)
				ui.SetActive(conf.Duration)
				squirters.Squirt(conf.Duration)
				timeout = time.Now().Add(conf.Cooldown + conf.Duration)
			}
		case <-ui.OnQuit():
			log.Println("Quitting!")
			return
		}
	}

}
