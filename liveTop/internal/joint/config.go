package joint

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	RequestLimit        int
	NATSurl             string
	ExpirationSeconds   int
	LoopIntervalSeconds int
}

func getDefaultConfig() *Config {
	return &Config{
		RequestLimit:        100,
		NATSurl:             "nats://nats:4222",
		ExpirationSeconds:   300,
		LoopIntervalSeconds: 1,
	}
}

func LoadConfigFromEnv() (*Config, error) {
	cfg := getDefaultConfig()

	err := setIntFromEnv("REQUEST_LIMIT", &cfg.RequestLimit)
	if err != nil {
		return &Config{}, err
	}

	err = setStringFromEnv("NATS_URL", &cfg.NATSurl)
	if err != nil {
		return &Config{}, err
	}

	err = setIntFromEnv("EXPIRATION_SECONDS", &cfg.ExpirationSeconds)
	if err != nil {
		return &Config{}, err
	}

	err = setIntFromEnv("LOOP_INTERVAL_SECONDS", &cfg.LoopIntervalSeconds)
	if err != nil {
		return &Config{}, err
	}

	return cfg, nil
}

func setStringFromEnv(key string, field *string) error {
	value, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}

	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}

	*field = value
	return nil
}

func setIntFromEnv(key string, field *int) error {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("invalid env var %s: %w", key, err)
	}

	if parsed <= 0 {
		return fmt.Errorf("invalid env var %s: must be positive", key)
	}

	*field = parsed
	return nil
}
