package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"Web3DataAnalyse/api"
)

var Cfg api.Config

func init() {
	configByte, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configByte, &Cfg)
	if err != nil {
		log.Fatal("Unmarshal yaml err", err)
		return
	}
}
