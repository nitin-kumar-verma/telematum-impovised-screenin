package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MySQLURI string
}

var GlobalConfig *AppConfig

func LoadConfig() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading envs")
	}

	uri, ok := os.LookupEnv("MYSQL_URI")
	if !(ok) {
		panic("MYSQL_URI missing")
	}
	GlobalConfig = &AppConfig{

		MySQLURI: uri,
	}
}
