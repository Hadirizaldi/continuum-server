package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func getEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	result, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return result
}

func getEnvInt64(key string, fallback int64) int64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	result, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return fallback
	}
	return result
}

func getEnvStringSlice(key string, fallback []string) []string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	parts := strings.Split(val, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	result, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}
	return result
}
