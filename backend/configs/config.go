package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type databaseConfig struct {
	Hostname     string
	Port         int
	Password     string
	Username     string
	DatabaseName string
}

type SURLBackendConfig struct {
	Database databaseConfig
	Port     int
}

var instance *SURLBackendConfig

func getValue(key string) string {
	return os.Getenv(key)
}

func getRequiredValue(key string) string {
	value := getValue(key)
	if value == "" {
		panic(fmt.Sprintf("required '%s' value not found in .env", key))
	}

	return value
}

func getRequiredValueToInt(key string) int {
	value := getRequiredValue(key)

	parsedInt, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("'%s' in .env should be an integer", key))
	}

	return parsedInt
}

func GetConfigs() *SURLBackendConfig {
	if instance == nil {
		godotenv.Load()
		instance = &SURLBackendConfig{
			Database: databaseConfig{
				Hostname:     getRequiredValue("DATABASE_HOSTNAME"),
				Username:     getRequiredValue("DATABASE_USERNAME"),
				Password:     getRequiredValue("DATABASE_PASSWORD"),
				DatabaseName: getRequiredValue("DATABASE_NAME"),
				Port:         getRequiredValueToInt("DATABASE_PORT"),
			},
			Port: getRequiredValueToInt("PORT"),
		}
	}

	return instance
}
