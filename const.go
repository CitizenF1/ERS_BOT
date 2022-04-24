package main

import tele "gopkg.in/telebot.v3"

var (
	mainMenu  = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnDetail = mainMenu.Text("Подробнее")
	btnProm   = mainMenu.Text("Связаться")
	btnC      = mainMenu.Contact("Contact")
	// Contact("Send Contanct")

	menuReserve = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnSch      = mainMenu.Text("Забронировать")
)

func InitMenu() {
	mainMenu.Reply(
		mainMenu.Row(btnDetail),
		mainMenu.Row(btnProm),
		mainMenu.Row(btnC),
	)
	menuReserve.Reply(
		menuReserve.Row(btnSch),
	)
}
