package main

import (
	"log"
	"os"
	"strconv"
)

// ConfigData contains config data. Api keys, paths
type ConfigData struct {
	botAPIKey     string
	dbPath        string
	removeTimeout int
}

func loadConfigFromEnv() ConfigData {

	botAPIKey, isSet := os.LookupEnv("TELEGRAM_BOT_API_KEY_VOICE")
	if !isSet {
		log.Panic("No bot api key found. Please set TELEGRAM_BOT_API_KEY_VOICE env")
	}

	dbPath, isSet := os.LookupEnv("TELEGRAM_BOT_DB_PATH")
	if !isSet {
		log.Panic("No bot db path found. Please set TELEGRAM_BOT_DB_PATH env")
	}

	timeout, isSet := os.LookupEnv("TELEGRAM_BOT_REMOVE_TIMEOUT")
	if !isSet {
		log.Panic("No timeout found. Please set TELEGRAM_BOT_DB_PATH env")
	}

	log.Println(botAPIKey + " " + dbPath + " " + timeout)

	timeoutInt, _ := strconv.Atoi(timeout)
	// add check later

	return ConfigData{botAPIKey, dbPath, timeoutInt}
}
