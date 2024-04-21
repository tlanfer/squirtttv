package main

import (
	"companion"
	"companion/adapter/inbound/streamlabs"
	"companion/adapter/inbound/twitchchat"
	"companion/adapter/inbound/yamlconfig"
	"companion/adapter/outbound/exchangerate"
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
	ui := trayicon.New()

	err := setupLogging()
	if err != nil {
		ui.ErrorMessage("%v", err.Error())
		return
	}

	time.Sleep(100 * time.Millisecond)
	log.Println("Companion starting...")

	loader := yamlconfig.New(filename, example)
	conf, err := loader.Load()
	if err != nil {
		if err == companion.ErrConfigNotFound {
			ui.ErrorMessage("No %v found, please create one. You can use %v as an example.", filename, example)
			ui.Quit()
			os.Exit(1)
		} else {
			ui.ErrorMessage("Failed to load config: %s", err.Error())
			ui.Quit()
			os.Exit(1)
		}
	}
	//conf.Dump(os.Stdout)

	converter, err := exchangerate.New(conf.Currency)
	if err != nil {
		ui.ErrorMessage("Currency %v can not be converted to: %v", conf.Currency, err)
		ui.Quit()
		os.Exit(1)
	}

	events := make(chan companion.StreamEvent)
	messages := make(chan companion.ChatMessage)

	if conf.Streamlabs != "" {
		sl := streamlabs.New(conf.Streamlabs, conf.Currency, converter)
		err := sl.Connect(events, messages)

		if err != nil {
			ui.ErrorMessage("Failed to connect to streamlabs: %s", err.Error())
		} else {
			ui.SetStreamlabsConnected(true)
		}
	}

	if conf.Twitch != "" {
		//tc := twitchchat.New(conf.Twitch, twitchchat.WithFdgt(), twitchchat.WithFaker(1*time.Second))
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
			squirters.Add(s)
		}
	} else {
		go squirters.Find()
	}

	timeout := time.Now()

	patternSquirt := func(p companion.SquirtPattern) {
		timeout = time.Now()
		for _, d := range p {
			timeout = timeout.Add(d)
		}

		for i, d := range p {
			if i%2 == 0 {
				ui.SetActive(d)
				squirters.Squirt(d)
			}
			time.Sleep(d)
		}
	}

	for {
		select {
		case m := <-messages:
			hasTrigger, pattern := conf.GetChatTrigger(m)
			if hasTrigger && time.Now().After(timeout) {
				log.Printf("Message from %v: %v -> Squirt for %v", m.User, m.Message, conf.Duration)
				go patternSquirt(*pattern)
			}
		case e := <-events:
			hasEvent, pattern := conf.GetEvent(e)
			if hasEvent && time.Now().After(timeout) {
				log.Printf("Got %v (%v): Squirt for %v", e.EventType, e.Amount, conf.Duration)
				go patternSquirt(*pattern)
			}
		case <-ui.OnQuit():
			log.Println("Quitting!")
			ui.Quit()
			return
		}
	}

}

func setupLogging() error {
	if os.Getenv("LOG_CONSOLE") == "" {

		file, err := os.OpenFile("companion.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}

		log.SetOutput(file)

	}

	return nil
}
