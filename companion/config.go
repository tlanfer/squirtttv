package companion

import (
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"strings"
	"time"
)

type ConfigError string

func (c ConfigError) Error() string {
	return string(c)
}

const ErrConfigNotFound ConfigError = "Config not found"

type ConfigLoader interface {
	Load() (*Config, error)
}

type Config struct {
	Cooldown time.Duration `yaml:"cooldown"`
	Duration SquirtPattern `yaml:"duration"`

	Twitch     string `yaml:"twitch"`
	Streamlabs string `yaml:"streamlabs"`
	Currency   string `yaml:"currency"`

	Squirters []string `yaml:"squirters"`

	Events       []Event       `yaml:"events"`
	ChatTriggers []ChatTrigger `yaml:"chat"`
}

type Event struct {
	Type    EventType     `yaml:"type"`
	Min     int           `yaml:"min"`
	Max     int           `yaml:"max,omitempty"`
	Pattern SquirtPattern `yaml:"duration,omitempty"`
}

type ChatTrigger struct {
	Role    ChatRole      `yaml:"role"`
	User    string        `yaml:"user,omitempty"`
	Message string        `yaml:"message"`
	Pattern SquirtPattern `yaml:"duration,omitempty"`
}

func (c Config) GetEvent(ev StreamEvent) (bool, *SquirtPattern) {
	defaultPattern := c.Duration

	var match *Event

	for i, e := range c.Events {

		if ev.EventType != e.Type {
			continue
		}

		if ev.Amount < e.Min {
			continue
		}

		if e.Max > 0 && ev.Amount > e.Max {
			continue
		}

		log.Println("found a match: ", e)

		if match == nil {
			match = &c.Events[i]
		}

		if e.Min > match.Min {
			match = &e
		}
	}

	if match != nil {
		if len(match.Pattern) == 0 {
			return true, &defaultPattern
		} else {
			return true, &match.Pattern
		}
	}

	return false, nil
}

func (c Config) GetChatTrigger(message ChatMessage) (bool, *SquirtPattern) {
	defaultPattern := c.Duration
	for _, e := range c.ChatTriggers {

		if message.Role < e.Role {
			continue
		}

		if !strings.Contains(message.Message, e.Message) {
			continue
		}

		if e.User != "" && message.User != e.User {
			continue
		}

		if len(e.Pattern) == 0 {
			return true, &defaultPattern
		} else {
			return true, &e.Pattern
		}
	}

	return false, nil
}

func (c Config) Dump(o io.Writer) {
	if c.Streamlabs != "" {
		c.Streamlabs = "REDACTED"
	}
	encoder := yaml.NewEncoder(o)
	encoder.SetIndent(2)
	encoder.Encode(c)
	return
}
