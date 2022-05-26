package utils

import (
	"log"
	"os"
)

type AppConfig struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
}

func LoadAppConfigFromEnv() *AppConfig {
	return &AppConfig{
		DbUser:     loadEnvByKey("DB_USER"),
		DbPassword: loadEnvByKey("DB_PASS"),
		DbPort:     loadEnvByKey("DB_PORT"),
		DbHost:     loadEnvByKey("DB_HOST"),
	}
}

func loadEnvByKey(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("can`t load env value by key=%s", key)
	}
	return val
}
