package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Web struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"web"`
	Telegram struct {
		BotToken          string `yaml:"botToken"`
		WitCiCreateChatID int    `yaml:"witCiCreateChatID"`
		BuildChatID       int    `yaml:"buildChatID"`
	} `yaml:"telegram"`
}

func readConfigFile(cfg *Config) {
	f, err := os.Open(currentDir() + "/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func readConfigEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
	}
}
