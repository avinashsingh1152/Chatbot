package providers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	DBConfig    DBConfig
	AppName     string
	AppPort     string
	GrpcSLMPort string
	S3          S3
}

type S3 struct {
	AwsRegion string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
}

func GetConfig() AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := AppConfig{
		AppName:     os.Getenv("APP_NAME"),
		AppPort:     os.Getenv("APP_PORT"),
		GrpcSLMPort: os.Getenv("GRPC_CLIENT_SLM_PORT"),
		S3: S3{
			AwsRegion: os.Getenv("AWS_REGION"),
		},
		DBConfig: DBConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			Port:       os.Getenv("DB_PORT"),
		},
	}

	return config
}
