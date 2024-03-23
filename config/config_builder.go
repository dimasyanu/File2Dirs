package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func InitConfig(configPath string) (AppSettings, error) {
	confFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
		return AppSettings{}, err
	}
	defer confFile.Close()

	confStr, err := io.ReadAll(confFile)
	if err != nil {
		log.Fatal(err)
		return AppSettings{}, err
	}

	config := Configuration{}
	err = json.Unmarshal(confStr, &config)
	if err != nil {
		log.Fatal(err)
		return AppSettings{}, err
	}

	return config.AppSettings, nil
}
