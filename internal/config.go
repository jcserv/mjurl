package internal

import "os"

type Configuration struct {
	Region      string
	Environment string
	Port        string
}

func NewConfiguration() (*Configuration, error) {
	cfg := &Configuration{}
	cfg.Region = os.Getenv("AWS_REGION")
	cfg.Environment = os.Getenv("ENVIRONMENT")
	cfg.Port = os.Getenv("PORT")
	return cfg, nil
}
