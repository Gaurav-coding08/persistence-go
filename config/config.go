package config

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type AppConfig struct {
	KafkaConfig KafkaConfig
	Port        string
	Env         string
	DBConfig    DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

type KafkaConfig struct {
	Broker   string
	Topic    string
	Username string
	Password string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		DBConfig: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Name:     getEnv("DB_NAME", "userdb"),
			Password: getEnv("DB_PASSWORD", "password"),
		},
		KafkaConfig: KafkaConfig{
			Broker:   getEnv("KAFKA_BROKER", "localhost:9092"),
			Topic:    getEnv("KAFKA_TOPIC", "stock_update"),
			Username: "user1",
			Password: "password-placeholder-0",
		},
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("APP_ENV", "local"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// func getKafkaPassword() string {
// 	cmd := exec.Command("sh", "-c", "kubectl get secret kafka-user-passwords --namespace default -o jsonpath='{.data.client-passwords}' | base64 -d | cut -d , -f 1")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		log.Fatalf("Failed to fetch Kafka password: %v", err)
// 	}
// 	return strings.TrimSpace(string(output))
// }

func getKafkaPassword() string {
	// First, try environment variable
	password := os.Getenv("KAFKA_PASSWORD")
	if password != "" {
		return password
	}

	// If running inside Kubernetes, fetch password from secret
	cmd := exec.Command("sh", "-c", "kubectl get secret kafka-user-passwords --namespace default -o jsonpath='{.data.client-passwords}' | base64 -d | cut -d , -f 1")
	output, err := cmd.Output()
	if err != nil {
		log.Println("⚠️ Failed to fetch Kafka password from Kubernetes Secret. Using default password.")
		return "default-password" // Change this to your actual default password
	}
	return strings.TrimSpace(string(output))
}
