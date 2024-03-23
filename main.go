package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"yanoo.id/file2dirs/config"
)

func main() {
	config, err := config.InitConfig("config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		Tidy(config.WatchDirectory)

		if config.IsOneTime {
			break
		}

		dt := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(dt, "|", "Standby..")
		time.Sleep(config.StandByDuration * time.Second)
	}
}

func Tidy(watchDir string) {
	files, err := os.ReadDir(watchDir)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Get file modified date
		stat, err := file.Info()
		if err != nil || stat.Name() == "desktop.ini" {
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		date := stat.ModTime().Format("2006_01_02")

		// If directory not exists, create it
		targetDir := filepath.Join(watchDir, date)
		_, err = os.Stat(targetDir)
		if err != nil && os.IsNotExist(err) {
			os.Mkdir(targetDir, os.ModePerm)
		}

		// Move target file
		currentLocation := filepath.Join(watchDir, stat.Name())
		targetLocation := filepath.Join(targetDir, stat.Name())
		err = os.Rename(currentLocation, targetLocation)
		if err != nil {
			log.Fatal(err)
			continue
		}

		fmt.Println("Moved:", currentLocation, "to", targetLocation)
	}
}
