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
		Host  string `yaml:"host"`
		Port  string `yaml:"port"`
		Debug bool   `yaml:"debug"`
	} `yaml:"web"`
	Tfs struct {
		ProjectUrl                           string `yaml:"projectUrl"`
		Pat                                  string `yaml:"pat"`
		AssessmentTasksQueue                 string `yaml:"assessmentTasksQueue"`
		AssessmentTasksDailyNotificationTime string `yaml:"assessmentTasksDailyNotificationTime"`
	} `yaml:"tfs"`
	Telegram struct {
		BotToken              string `yaml:"botToken"`
		PRCreatesChatID       int    `yaml:"prCreatesChatID"`
		CRCreatesChatID       int    `yaml:"crCreatesChatID"`
		CIUpdatesChatID       int    `yaml:"ciUpdatesChatID"`
		CICreatesChatID       int    `yaml:"ciCreatesChatID"`
		BuildChatID           int    `yaml:"buildChatID"`
		DeployChatID          int    `yaml:"deployChatID"`
		AssessmentTasksChatID int    `yaml:"assessmentTasksChatID"`
	} `yaml:"telegram"`
	ExcludedPipelines struct {
		Build   []string `yaml:"build"`
		Release []string `yaml:"release"`
	} `yaml:"excludedPipelines"`
}

func readConfigFile(cfg *Config) {
	f, err := os.Open("./config/config.yml")
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
