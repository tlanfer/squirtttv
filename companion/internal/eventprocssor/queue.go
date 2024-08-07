package eventprocssor

import (
	"companion/internal/adapter/outbound/squirter"
	"companion/internal/config"
	"time"
)

type queuedAction struct {
	p       config.SquirtPattern
	devices []squirter.Squirter
}

var queue = make(chan queuedAction, 1000)
var stop = make(chan struct{})

func (p *processor) run() {
	for {
		select {
		case action := <-queue:
			if len(action.devices) > 0 {
				sendPattern(action.p, action.devices)
				time.Sleep(time.Duration(p.config.Settings.SprayPause))
			}
		case <-stop:
			return
		}
	}
}

func AddToQueue(p config.SquirtPattern, devices []squirter.Squirter) {
	queue <- queuedAction{
		p:       p,
		devices: devices,
	}
}
