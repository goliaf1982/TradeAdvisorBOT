package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string

    BinanceAPIKey    string
    BinanceSecretKey string
}

func LoadConfig() Config {
    err := godotenv.Load("config/.env")
    if err != nil {
	log.Fatalf("Не вдалося завантажити .env файл: %v", err)
    }

    return Config{
	DBUser:           os.Getenv("DB_USER"),
	DBPassword:       os.Getenv("DB_PASSWORD"),
	DBName:           os.Getenv("DB_NAME"),
	DBPort:           os.Getenv("DB_PORT"),
	BinanceAPIKey:    os.Getenv("BINANCE_API_KEY"),
	BinanceSecretKey: os.Getenv("BINANCE_SECRET_KEY"),
    }
}
