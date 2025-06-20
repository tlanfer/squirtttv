package config

import (
	"companion/internal"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Config struct {
	Settings Settings `json:"settings" yaml:"settings"`
	Devices  []Device `json:"devices" yaml:"devices"`
	Events   Events   `json:"events" yaml:"events"`
}

type Settings struct {
	Twitch         string   `json:"twitch" yaml:"twitch"`
	Streamlabs     string   `json:"streamlabs" yaml:"streamlabs"`
	Streamelements string   `json:"streamelements" yaml:"streamelements"`
	BaseCurrency   string   `json:"baseCurrency" yaml:"baseCurrency"`
	SprayPause     Duration `json:"sprayPause" yaml:"sprayPause"`
	GlobalCooldown Duration `json:"globalCooldown" yaml:"globalCooldown"`
}

type Device struct {
	Name string `json:"name" yaml:"Name"`
	Host string `json:"host" yaml:"Host"`
}

type Events struct {
	Bits    []Event `json:"bits" yaml:"bits"`
	Dono    []Event `json:"dono" yaml:"dono"`
	Resubt1 []Event `json:"resubt1" yaml:"resubt1"`
	Resubt2 []Event `json:"resubt2" yaml:"resubt2"`
	Resubt3 []Event `json:"resubt3" yaml:"resubt3"`
	Gifts   []Event `json:"gifts" yaml:"gifts"`

	ChatMessage []ChatMessage `json:"chat" yaml:"chat"`
}

type Event struct {
	Amount  float64       `json:"amount" yaml:"amount"`
	Match   string        `json:"match" yaml:"match"`
	Pattern SquirtPattern `json:"pattern" yaml:"pattern"`
	Choose  string        `json:"choose" yaml:"choose"`
	Devices []string      `json:"devices" yaml:"devices"`
}

type ChatMessage struct {
	Trigger string              `json:"trigger" yaml:"trigger"`
	Roles   []internal.ChatRole `json:"roles" yaml:"roles"`
	Pattern SquirtPattern       `json:"pattern" yaml:"pattern"`
	Choose  string              `json:"choose" yaml:"choose"`
	Devices []string            `json:"devices" yaml:"devices"`
}

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func (e Event) String() string {
	return fmt.Sprintf("Event{Amount: %f, Match: %s, Pattern: %v, Choose: %s, Devices: %v}", e.Amount, e.Match, e.Pattern, e.Choose, e.Devices)
}
