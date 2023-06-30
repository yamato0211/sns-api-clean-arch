package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	HTTPInfo     *HTTPInfo
	PostgresInfo *PostgresInfo
}

type HTTPInfo struct {
	Addr string
}

type PostgresInfo struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresDB       string
	PostgresPort     string
}

func LoadConfig() *appConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env not found")
	}

	addr := ":" + os.Getenv("API_PORT")

	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPass := os.Getenv("POSTGRES_PASSWORD")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresPort := os.Getenv("POSTGRES_PORT")

	dbInfo := &PostgresInfo{
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPass,
		PostgresHost:     postgresHost,
		PostgresDB:       postgresDB,
		PostgresPort:     postgresPort,
	}

	conf := appConfig{
		HTTPInfo:     httpInfo,
		PostgresInfo: dbInfo,
	}

	return &conf
}
