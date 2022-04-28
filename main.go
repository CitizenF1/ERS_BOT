package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
	ctx, cancel := context.WithCancel(context.Background())
	dbCancel = cancel
	initDB(ctx)
}

func main() {
	defer dbCancel()
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
	InitMenuButtons()
	setHandlers()

	botInstance.Start()
}
