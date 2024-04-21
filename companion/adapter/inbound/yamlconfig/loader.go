package yamlconfig

import (
	"companion"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"time"
)

func New(filename, example string) companion.ConfigLoader {
	return &loader{filename: filename, example: example}
}

type loader struct {
	filename string
	example  string
}

func (l *loader) Load() (*companion.Config, error) {
	l.Example()

	file, err := os.OpenFile(l.filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, companion.ErrConfigNotFound
		} else {
			return nil, err
		}
	}
	defer file.Close()

	c := &companion.Config{
		AllowLegacy: false,
	}
	yaml.NewDecoder(file).Decode(c)

	if c.Twitch == "" && c.Streamlabs == "" {
		return nil, errors.New("need either a twitch channel or a streamlabs token. see config.example.yaml for an example")
	}

	if len(c.Duration) == 0 {
		c.Duration = []time.Duration{1 * time.Second}
	}

	for _, duration := range c.Duration {
		if duration < 500*time.Millisecond {
			duration = 500 * time.Millisecond
		}
	}

	if len(c.Events) == 0 && len(c.ChatTriggers) == 0 {
		return nil, errors.New("must give at least one event to trigger on")
	}

	if c.Currency == "" {
		c.Currency = "EUR"
	}
	c.Currency = strings.ToLower(c.Currency)

	return c, nil
}

func (l *loader) Example() {
	file, err := os.OpenFile(l.example, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	c := &companion.Config{
		Cooldown: 5 * time.Second,
		Duration: []time.Duration{1 * time.Second},
		Squirters: []string{
			"192.168.1.200",
		},
		Twitch:     "tlanfer",
		Streamlabs: "eyJ0eX.... get yours from https://streamlabs.com/dashboard#/settings/api-settings > API Tokens > Your Socket API Token",
		Events: []companion.Event{
			{Type: companion.EventTypeBits, Min: 0, Max: 100},
			{Type: companion.EventTypeBits, Min: 200, Max: 250},
			{Type: companion.EventTypeDono, Min: 20, Max: 30},
			{Type: companion.EventTypeDono, Min: 100},
			{Type: companion.EventTypeSub, Min: 10},
			{Type: companion.EventTypeSub, Min: 25, Pattern: companion.SquirtPattern{
				3 * time.Second,
			}},
			{Type: companion.EventTypeSub, Min: 50, Pattern: companion.SquirtPattern{
				1 * time.Second,
				500 * time.Millisecond,
				2 * time.Second,
				500 * time.Millisecond,
				3 * time.Second,
			}},
		},
		ChatTriggers: []companion.ChatTrigger{
			{
				Role:    companion.ChatRoleMod,
				Message: "!squirt",
			},
		},
	}
	yaml.NewEncoder(file).Encode(c)
}
