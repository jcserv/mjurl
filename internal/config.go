package internal

import "os"

type Configuration struct {
	Region      string
	Environment string
	HTTPPort    string
	GRPCPort    string
	DatabaseURL string
}

func NewConfiguration() (*Configuration, error) {
	cfg := &Configuration{}
	cfg.Region = GetString("AWS_REGION", "us-east-1")
	cfg.Environment = GetString("ENVIRONMENT", "prod")
	cfg.HTTPPort = GetString("HTTP_PORT", "8080")
	cfg.GRPCPort = GetString("GRPC_PORT", "50051")
	cfg.DatabaseURL = GetString("DATABASE_URL", "postgres://admin:123@localhost:5432/postgres")
	return cfg, nil
}

func GetString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
