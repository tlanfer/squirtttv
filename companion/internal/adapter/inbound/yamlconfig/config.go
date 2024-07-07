package yamlconfig

import (
	"companion/internal/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Init(filename string) {
	l := loader{filename: filename}
	l.Load()
	config.Subscribe(func(c config.Config) {
		l.Write(c)
	})
}

type loader struct {
	filename string
}

func (l *loader) Load() {
	file, err := os.OpenFile(l.filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println("config file not found")
		return
	}
	defer file.Close()

	c := config.Config{}
	if err := yaml.NewDecoder(file).Decode(&c); err != nil {
		log.Println("error decoding config file:", err)
		return
	}

	config.Set(c)

	return
}

func (l *loader) Write(c config.Config) {
	file, err := os.OpenFile(l.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	if err := yaml.NewEncoder(file).Encode(c); err != nil {
		return
	}
}
