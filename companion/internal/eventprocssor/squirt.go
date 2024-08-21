package eventprocssor

import (
	"companion/internal/adapter/outbound/squirter"
	"companion/internal/config"
	"log"
	"math/rand"
	"time"
)

func squirt(ev config.Event) {
	var hosts []string

	switch ev.Choose {
	case "all":
		dv := config.Get().Devices
		for _, device := range dv {
			hosts = append(hosts, device.Host)
		}

	case "allOf":
		hosts = ev.Devices

	case "oneOf":
		if len(ev.Devices) == 0 {
			log.Println("No devices to squirt for event", ev)
		} else {
			hosts = []string{ev.Devices[rand.Intn(len(ev.Devices))]}
		}
	}

	var squirters []squirter.Squirter
	for _, host := range hosts {
		squirters = append(squirters, squirter.Squirter{Host: host})
	}

	AddToQueue(ev.Pattern, squirters)
}

func sendPattern(p config.SquirtPattern, devices []squirter.Squirter) {
	for i, d := range p {
		if i%2 == 0 {
			send(d, devices)
		}
		time.Sleep(d)
	}
}

func send(duration time.Duration, devices []squirter.Squirter) {
	for _, d := range devices {
		go d.Squirt(duration)
	}
}
