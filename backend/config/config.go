package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var SecretKeys Secrets

type Secrets struct {
	SessionSecret string `yaml:"jwt_secret"`
	PepperPas     string `yaml:"pepper_pas"`
	Mysql_pas     string `yaml:"mysql_pas"`
	Redis_pas     string `yaml:"redis_pas"`
}

type PepperPas struct {
}

func (c *Secrets) parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func GetSecret() {
	data, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	var config Secrets
	if err := config.parse(data); err != nil {
		log.Fatal(err)
	}
	SecretKeys = config
}
