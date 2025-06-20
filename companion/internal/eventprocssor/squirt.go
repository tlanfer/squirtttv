package eventprocssor

import (
	"companion/internal/adapter/outbound/squirter"
	"companion/internal/config"
	"log"
	"math/rand"
	"time"
)

func squirt(choose string, devices []string, pattern config.SquirtPattern) {
	var hosts []string

	switch choose {
	case "all":
		dv := config.Get().Devices
		for _, device := range dv {
			hosts = append(hosts, device.Host)
		}

	case "allOf":
		hosts = devices

	case "oneOf":
		if len(devices) == 0 {
			log.Println("No devices to squirt for event")
		} else {
			hosts = []string{devices[rand.Intn(len(devices))]}
		}
	}

	var squirters []squirter.Squirter
	for _, host := range hosts {
		squirters = append(squirters, squirter.Squirter{Host: host})
	}

	AddToQueue(pattern, squirters)
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
