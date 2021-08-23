package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server
	Mongo
}

type Server struct {
	Mode string
	Port string
}

type Mongo struct {
	Username string
	Password string
	Uri      string
	Database string
}

var AppConfig = &Config{}

var ServerConfig = &AppConfig.Server
var MongoConfig = &AppConfig.Mongo

func Setup() {
	yamlFile, err := ioutil.ReadFile("config/environment.dev.yaml")
	if err != nil {
		panic("Error...while reading the yaml file...")
	}
	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		panic("Error...while unmarshal yaml file...")
	}
}
