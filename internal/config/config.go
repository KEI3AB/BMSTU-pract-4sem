package config

import (
	"os"
	"strconv"
)

type Config struct {
	BalancerPort       string
	WorkerPort         string
	Strategy           string
	LoadGenTarget      string
	LoadGenConcurrency int
	LoadGenRequests    int
}

func Load() *Config {
	return &Config{
		BalancerPort:       getEnv("BALANCER_PORT", "8080"),
		WorkerPort:         getEnv("WORKER_PORT", "8000"),
		Strategy:           getEnv("STRATEGY", "LEAST_CONN"),
		LoadGenTarget:      getEnv("LOADGEN_TARGET", "http://localhost:8080"),
		LoadGenConcurrency: getEnvAsInt("LOADGEN_CONCURRENCY", 50),
		LoadGenRequests:    getEnvAsInt("LOADGEN_REQUESTS", 500),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return fallback
}
