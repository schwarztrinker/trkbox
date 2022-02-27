package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf ServerConf

type ServerConf struct {
	JwtSecret   string `yaml:"jwtSecret"`
	Port        string `yaml:"port"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
	DB_USER     string `yaml:"DB_USER"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_PORT     int    `yaml:"DB_PORT"`
}

func (c *ServerConf) GetConf() *ServerConf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
