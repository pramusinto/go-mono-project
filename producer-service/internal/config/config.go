package config

import "os"

type Config struct {
	KafkaBroker string
	KafkaTopic  string
}

func LoadConfig() *Config {
	return &Config{
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
		KafkaTopic:  getEnv("KAFKA_TOPIC", "events"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
