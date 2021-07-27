package config

type Config struct {
	ConnectionString string `envconfig:"CONNECTION_STRING"`
	SlackToken       string `envconfig:"SLACK_TOKEN" required:"TRUE"`
	FrontEndURL      string `envconfig:"FRONTEND_URL" required:"TRUE"`
}

// Load parses the env varaibles into a config struct
func Load() (*Config, error) {
	cfg := Config{}

	// err := envconfig.Process("", &cfg)
	// if err != nil {
	// 	return nil, fmt.Errorf("error processing env config: %w", err)
	// }
	return &cfg, nil
}
