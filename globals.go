package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BotToken              string
	envListenPort         string
	Poller                string
	rootTelegramMethodURL string
	botMode               string
	publicURL             string
	mongoHostname         string
	mongoPort             string
)

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultVal
}

func init() {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	BotToken = os.Getenv("BOT_TOKEN")
	rootTelegramMethodURL = "https://api.telegram.org/bot" + BotToken

	// envRootPublicURL = os.Getenv("ROOT_PUBLIC_URL")
	// publicURL = envRootPublicURL + "/" + envBotToken

	envListenPort = getEnv("LISTEN_PORT", "9000")
	botMode = getEnv("BOT_MODE", "development")

	mongoHostname = getEnv("MONGO_HOSTNAME", "localhost")
	mongoPort = getEnv("MONGO_PORT", "27017")
}
