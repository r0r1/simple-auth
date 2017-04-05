package configs

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DB_HOST          string
	DB_NAME          string
	DB_USER          string
	DB_PASSWORD      string
	APP_ID           string
	APP_SECRET       string
	APP_CALLBACK_URL string
	APP_URL          string
}

// Read Config
func ReadConfig() Config {
	var configfile = "./configs/config"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	log.Print(config.DB_NAME)
	return config
}
