package main

import (
	"log"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file", err)
	}

	pref := getBotPref()
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	botInstance = b
	log.Println("Bot Started!")
	InitMenu()
	setHandlers()
	botInstance.Start()
}
