package config

type Config struct {
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
