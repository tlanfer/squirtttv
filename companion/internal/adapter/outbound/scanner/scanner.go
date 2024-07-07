package scanner

import (
	"companion/internal/adapter/outbound/squirter"
	"context"
	"github.com/grandcat/zeroconf"
	"log"
	"time"
)

var allDevices = make(map[string]bool)

func List() []string {
	var devices []string
	for k := range allDevices {
		devices = append(devices, k)
	}
	return devices
}

func New() {
	go scan()

	go func() {
		t := time.NewTicker(1 * time.Minute)
		for range t.C {
			scan()
		}
	}()
}

func scan() {

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
			host := entry.AddrIPv4[0].String()

			s := squirter.Squirter{Host: host}
			if s.IsSquirter() && !allDevices[host] {
				log.Println("Found squirter:", host)
				allDevices[host] = true
			}
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
