package main

import (
	tele "gopkg.in/telebot.v3"
)

var botInstance *tele.Bot

func getBotPref() tele.Settings {
	// var poller tele.Poller = &tele.Webhook{
	// 	Listen:   listenPort,
	// 	Endpoint: &tele.WebhookEndpoint{PublicURL: publicURL},
	// }
	pref := tele.Settings{
		Token: BotToken,
		// Poller: poller,
	}

	return pref
}
