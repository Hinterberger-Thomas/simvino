package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	RedisConf struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
		Db   uint8  `yaml:"dbname"`
	} `yaml:"redis"`

	PostgresConf struct {
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"postgres"`

	Mongo struct {
		Host string `yaml:"host"`
	} `yaml:"mongo"`

	Password struct {
		Pepper string `yaml:"pepper"`
	} `yaml:"password"`
}

var Configs Conf

func InitConf() {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Configs)
	if err != nil {
		log.Fatal(err)
	}

}
