package internal

import "os"

type Configuration struct {
	Region      string
	Environment string
	Port        string
	DatabaseURL string
}

func NewConfiguration() (*Configuration, error) {
	cfg := &Configuration{}
	cfg.Region = GetString("AWS_REGION", "us-east-1")
	cfg.Environment = GetString("ENVIRONMENT", "prod")
	cfg.Port = GetString("PORT", "8080")
	cfg.DatabaseURL = GetString("DATABASE_URL", "postgres://admin:123@localhost:5432/postgres")
	return cfg, nil
}

func GetString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
