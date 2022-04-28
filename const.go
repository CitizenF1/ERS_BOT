package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	tele "gopkg.in/telebot.v3"
)

var (
	mainMenu  = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnDetail = mainMenu.Text("Подробнее")
	// btnProm   = mainMenu.Text("Связаться")
	btnContact = mainMenu.Contact("Связаться")
	// Contact("Send Contanct")

	menuReserve = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnSch      = mainMenu.Text("Забронировать")
)

const (
	//CallBacks for back office
	CallbackManager1 = "manager1_call"
	CallbackManager2 = "manager2_call"
	CallbackManager3 = "manager3_call"
)

type StoredUsers struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	User *tele.User         `bson:"user"`
}

func InitMenuButtons() {
	mainMenu.Reply(
		mainMenu.Row(btnDetail),
		// mainMenu.Row(btnProm),
		mainMenu.Row(btnContact),
	)
	menuReserve.Reply(
		menuReserve.Row(btnSch),
	)
}
