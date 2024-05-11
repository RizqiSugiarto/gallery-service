package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type HttpServer struct{
	Addr string
}

type GrpcServer struct {
	Addr string 
}

type Database struct {
	UrlDb string
}

type Config struct {
	HttpServer
	GrpcServer
	Database
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Config - Load - godotenv.Load: %v", err)
	}

	cfg := &Config{
		HttpServer: HttpServer{
			Addr: os.Getenv("HTTP_SERVER_ADDR"),
		},
		GrpcServer: GrpcServer{
			Addr: os.Getenv("GRPC_SERVER_ADDR"),
		},
		Database: Database{
			UrlDb: os.Getenv("URL_DB"),
		},
	}
	return cfg
}