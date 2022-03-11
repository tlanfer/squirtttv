package companion

import "time"

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
	Duration time.Duration `yaml:"duration"`

	Twitch     string `yaml:"twitch"`
	Streamlabs string `yaml:"streamlabs"`

	Squirters []string `yaml:"squirters"`

	Events []Event `yaml:"events"`
}

type Event struct {
	Type EventType `yaml:"type"`
	Min  int       `yaml:"min"`
	Max  int       `yaml:"max,omitempty"`
}

func (c Config) Matches(ev StreamEvent) bool {
	for _, e := range c.Events {

		if ev.EventType != e.Type {
			continue
		}

		if ev.Amount < e.Min {
			continue
		}

		if e.Max > 0 && ev.Amount > e.Max {
			continue
		}

		return true
	}

	return false
}
