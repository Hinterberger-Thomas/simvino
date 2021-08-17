package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type sessionSecret struct {
	sessionSecret string `yaml:"session_secret"`
}

func (c *sessionSecret) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func GetSessionSecret() string {
	data, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	var config sessionSecret
	if err := config.Parse(data); err != nil {
		log.Fatal(err)
	}
	return config.sessionSecret
}
