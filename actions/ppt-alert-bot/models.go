package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type TGConfig struct {
	TelegramAPI struct {
		APIToken string `yaml:"api-token"`
		ChatId   int64  `yaml:"chat-id"`
	} `yaml:"telegram-api"`
}

func NewConfig(path string) (*TGConfig, error) {
	config := &TGConfig{}
	if file, err := os.Open(path); err != nil {
		return nil, err
	} else {
		defer file.Close()
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			return nil, err
		} else {
			return config, nil
		}
	}
}
