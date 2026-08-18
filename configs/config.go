package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/disgoorg/snowflake/v2"
	"github.com/joho/godotenv"
)

var (
	botRequired     = ""
	backendRequired = ""
)

type databaseConfig struct {
	Hostname     string
	Port         int
	Password     string
	Username     string
	DatabaseName string
}

type botConfig struct {
	Token                 string
	OwnerID               snowflake.ID
	BackendURL            string
	DevOnlyCommandGuildID snowflake.ID
}

type backendConfig struct {
	Port int
}

type SURLConfig struct {
	Backend  backendConfig
	Bot      botConfig
	Database databaseConfig
}

var instance *SURLConfig
var once sync.Once

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

func GetConfig() *SURLConfig {
	once.Do(func() {
		_ = godotenv.Load()
		instance = &SURLConfig{
			Database: databaseConfig{
				Hostname:     getRequiredValue("DATABASE_HOSTNAME"),
				Username:     getRequiredValue("DATABASE_USERNAME"),
				Password:     getRequiredValue("DATABASE_PASSWORD"),
				DatabaseName: getRequiredValue("DATABASE_NAME"),
				Port:         getRequiredValueToInt("DATABASE_PORT"),
			},
		}

		if backendRequired == "true" {
			instance.Backend = backendConfig{
				Port: getRequiredValueToInt("BACKEND_PORT"),
			}
		}

		if botRequired == "true" {
			instance.Bot = botConfig{
				Token:      getRequiredValue("BOT_TOKEN"),
				OwnerID:    snowflake.MustParse(getRequiredValue("BOT_OWNER_ID")),
				BackendURL: getRequiredValue("BOT_BACKEND_URL"),
			}
		}
	})

	return instance
}
