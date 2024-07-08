package yamlconfig

import (
	"companion/internal/adapter/inbound/trayicon"
	"companion/internal/config"
	"github.com/sqweek/dialog"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
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
		config.Set(config.Config{
			Settings: config.Settings{
				BaseCurrency:   "eur",
				SprayPause:     config.Duration(3 * time.Second),
				GlobalCooldown: config.Duration(0 * time.Second),
			},
		})
		if dialog.Message("No config file found. Open UI to configure things?").YesNo() {
			trayicon.OpenUI()
		}
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("error closing config file:", err)
		}
	}()

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

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("error closing config file:", err)
		}
	}()

	if err := yaml.NewEncoder(file).Encode(c); err != nil {
		return
	}
}
