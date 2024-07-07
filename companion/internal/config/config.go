package config

type Subscriber func(c Config)

var config Config
var subscribers []Subscriber

func Subscribe(s Subscriber) {
	subscribers = append(subscribers, s)
}

func Get() Config {
	return config
}

func Set(c Config) {
	config = c
	for _, s := range subscribers {
		go s(c)
	}
}
