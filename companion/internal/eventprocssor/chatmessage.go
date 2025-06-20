package eventprocssor

import (
	"companion/internal"
	"companion/internal/config"
	"log"
	"strings"
)

func (p *processor) processMessage(msg internal.ChatMessage) {
	c := config.Get()
	events := c.Events.ChatMessage

	for _, event := range events {
		if !strings.Contains(msg.Message, event.Trigger) {
			continue
		}

		if !isRoleMatch(msg.Role, event.Roles) {
			log.Printf("Skipping event for user %s with role %s, required roles: %v", msg.User, msg.Role, event.Roles)
			continue
		}

		squirt(event.Choose, event.Devices, event.Pattern)
		return
	}
}
