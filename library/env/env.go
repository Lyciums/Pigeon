package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func WriteByKey(key, value string) {
	em, _ := godotenv.Unmarshal(key + "=" + value)
	WriteByMap(em)
}

func WriteByMap(em map[string]string) {
	_ = godotenv.Write(em, ".env")
}

func init() {
	_ = godotenv.Load()
}
