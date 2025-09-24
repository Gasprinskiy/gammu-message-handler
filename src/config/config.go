package config

import (
	"fmt"
	"os"
)

type Config struct {
	BotToken       string
	ChatID         string
	GrpcPort       string
	HttpServerPort string
	DbName         string
	DbPassword     string
	DbUser         string
	DbHost         string
	DbPort         string
}

func (c Config) DbConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}

func NewConfig() *Config {
	return &Config{
		BotToken:       os.Getenv("BOT_TOKEN"),
		ChatID:         os.Getenv("CHAT_ID"),
		HttpServerPort: fmt.Sprintf(":%s", os.Getenv("HTTP_SERVER_PORT")),
		DbName:         os.Getenv("DB_NAME"),
		DbPassword:     os.Getenv("DB_PASSWORD"),
		DbUser:         os.Getenv("DB_USER"),
		DbHost:         os.Getenv("DB_HOST"),
		DbPort:         os.Getenv("DB_PORT"),
	}
}
