package cfg

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Global struct {
	App  Application
	Aut  Auth   `yaml:"auth"`
	My   MySQL  `yaml:"mysql"`
	Serv Server `yaml:"server"`
}

type Application struct {
	TLoc *time.Location
}

type Auth struct {
	Key string `yaml:"key"`
}

type MySQL struct {
	Dialect string `yaml:"dialect"`
	DSN     string `yaml:"dsn"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var (
	Glb  Global
	App  Application
	Aut  Auth
	My   MySQL
	Serv Server
)

func Load(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Println("error: config file failed to read - ", err)
		return err
	}

	err = yaml.Unmarshal(data, &Glb)
	if err != nil {
		log.Println("error: config file failed to unmarshal - ", err)
		return err
	}

	App = Glb.App
	Aut = Glb.Aut
	My = Glb.My
	Serv = Glb.Serv

	App.TLoc, _ = time.LoadLocation(AppTimeZone)

	return nil
}
