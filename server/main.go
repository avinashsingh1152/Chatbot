package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"server/client"
	"server/core"
	"server/httpServer"
	"server/providers"
	"server/repository"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {

	config := providers.GetConfig()
	db, err := setupDatabase(config)
	if err != nil {
		log.Printf("Could not setup database: %v", err)
		panic(err)
	}

	S3Config, err := NewS3Instance(config)
	if err != nil {
		log.Printf("error while connection with AWS: %v", err)
		panic(err)
	}

	apiRouter := mux.NewRouter()
	http.Handle("/", apiRouter)

	httpLogger := log.New(os.Stdout, "http Server: ", log.LstdFlags)
	coreLogger := log.New(os.Stdout, "core Layer: ", log.LstdFlags)
	repoLogger := log.New(os.Stdout, "DB Layer: ", log.LstdFlags)
	grpcLogger := log.New(os.Stdout, "grpc Layer: ", log.LstdFlags)
	server := httpServer.HttpServer{
		Router: apiRouter,
		Core: core.Core{
			Logger: coreLogger,
			S3:     S3Config,
			SlModelGrpcClient: client.SlModelGrpcClient{
				AppConfig: config,
				Logger:    grpcLogger,
			},
			DB: repository.SqlRepository{
				Logger: repoLogger,
				DB:     db,
			},
		},
		Logger: httpLogger,
	}
	server.Init()

	if err := http.ListenAndServe(config.AppPort, nil); err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}

func setupDatabase(config providers.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBConfig.DBUser, config.DBConfig.DBPassword, config.DBConfig.DBHost, config.DBConfig.Port, config.DBConfig.DBName)

	// Open the database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewS3Instance(config providers.AppConfig) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.S3.AwsRegion), // Set your AWS region
	})

	cong := aws.NewConfig().WithRegion(config.S3.AwsRegion)
	s3Instance := s3.New(sess, cong)
	if err != nil {
		return nil, err
	}

	return s3Instance, nil
}
