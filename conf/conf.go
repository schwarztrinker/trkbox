package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf ServerConf

type ServerConf struct {
	JwtSecret string `yaml:"jwtSecret"`
	Port      string `yaml:"port"`
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
