package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"
)

type Config struct {
	Environment string

	ServerHost string
	HTTPPort   string

	PostgresHost          string
	PostgresUser          string
	PostgresDatabase      string
	PostgresPassword      string
	PostgresPort          int
	PostgresMaxConnection int32

	DefaultOffset int
	DefaultLimit  int

	DefaultFrom int
	DefaultTo   int
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.DefaultFrom = 0
	cfg.DefaultTo = 1000000000

	cfg.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", ReleaseMode))

	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVER_HOST", "localhost"))
	cfg.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "abdulbosit"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "exam"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "946236953"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))

	cfg.PostgresMaxConnection = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_PORT", 30))

	cfg.DefaultOffset = cast.ToInt(getOrReturnDefaultValue("OFFSET", 0))
	cfg.DefaultLimit = cast.ToInt(getOrReturnDefaultValue("LIMIT", 10))

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
