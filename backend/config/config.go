package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Postgres Postgres
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() Config {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println("NO .env file . file not found")
	}
	cfg := Config{
		Postgres{
			Host:     cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost")),
			Port:     cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT", "5432")),
			Database: cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "blog_post")),
			User:     cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "admin")),
			Password: cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "admin123")),
		},
	}

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
