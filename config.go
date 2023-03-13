package main

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	User struct {
		Username string `yaml:"username"`
		PIva     string `yaml:"piva"`
		Pincode  string `yaml:"pincode"`
	} `yaml:"user"`
}

func getConfig() (*Config, error) {
	file, err := os.Open("config.yml")
	if err != nil {
		log.Printf("Failed to open file: %s\n", err)
		return nil, errors.New("failed to open config file")
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to decode config file")
	}

	return &cfg, nil
}
