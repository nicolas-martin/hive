package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ConnectionString string `envconfig:"CONNECTION_STRING"`
	SlackToken       string `envconfig:"SLACK_TOKEN" required:"TRUE"`
	FrontEndURL      string `envconfig:"FRONTEND_URL" required:"TRUE"`
}

// Load parses the env varaibles into a config struct
func Load() *Config {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Panic("error procesing env configs")
	}
	return &cfg
}
