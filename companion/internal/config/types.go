package config

type Config struct {
	Settings Settings `json:"settings" yaml:"settings"`
	Devices  []Device `json:"devices" yaml:"devices"`
	Events   Events   `json:"events" yaml:"events"`
}

type Settings struct {
	Twitch         string `json:"twitch" yaml:"twitch"`
	Streamlabs     string `json:"streamlabs" yaml:"streamlabs"`
	Streamelements string `json:"streamelements" yaml:"streamelements"`
	BaseCurrency   string `json:"baseCurrency" yaml:"baseCurrency"`
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
}

type Event struct {
	Amount  int           `json:"amount" yaml:"amount"`
	Match   string        `json:"match" yaml:"match"`
	Pattern SquirtPattern `json:"pattern" yaml:"pattern"`
	Choose  string        `json:"choose" yaml:"choose"`
	Devices []string      `json:"devices" yaml:"devices"`
}
