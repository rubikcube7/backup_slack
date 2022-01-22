package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	ApiToken     string
	ApiUserID    string
	ApiChannelID string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiToken:     cfg.Section("slack").Key("api_token").String(),
		ApiUserID:    cfg.Section("slack").Key("api_user").String(),
		ApiChannelID: cfg.Section("slack").Key("api_channel").String(),
	}
}
