package test

import (
	"companion/internal/adapter/outbound/squirter"
	"companion/internal/config"
	"companion/internal/eventprocssor"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		d := dto{}
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var devices []string

		switch d.Choose {
		case "all":
			dv := config.Get().Devices
			for _, device := range dv {
				devices = append(devices, device.Host)
			}

		case "allOf":
			devices = d.Devices

		case "oneOf":
			if len(d.Devices) > 0 {
				devices = []string{
					d.Devices[rand.Intn(len(d.Devices))],
				}
			}
		}

		if len(devices) == 0 {
			http.Error(w, "No devices selected", http.StatusBadRequest)
			return
		}

		var squirters []squirter.Squirter
		for _, host := range devices {
			h := host
			squirters = append(squirters, squirter.Squirter{Host: h})
		}

		log.Println("Adding to queue", d.Pattern, squirters)
		eventprocssor.AddToQueue(d.Pattern, squirters)
	})
}

type dto struct {
	Pattern config.SquirtPattern `json:"pattern"`
	Choose  string               `json:"choose"`
	Devices []string             `json:"devices"`
}
