package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	tele "gopkg.in/telebot.v3"
)

func setHandlers() {
	botInstance.Handle("/start", startHandler)

	botInstance.Handle(&btnDetail, DetailHandler)
	botInstance.Handle(&btnSch, ReserveHandler)
	botInstance.Handle(&btnProm, ConnecHandler)
}

func startHandler(c tele.Context) error {
	WriteMaster(*c.Sender())
	return c.Send("Вас приветствует WANT.INVEST!\nУ нас для вас прекрасная новость: мы запускаем четвёртый поток марафона «Криптотрейдинг с нуля».\nСтоимость: 50 000 тенге\nСтарт: 29 апреля\n\nЕсли вы сегодня сделаете предоплату, цена для вас составит 38 900 тенге 🎁\n\nВам интересно наше предложение?", mainMenu)
}

func DetailHandler(c tele.Context) error {
	// var buttonArray [][]tele.InlineButton
	p := &tele.Photo{File: tele.FromDisk("./img/photo_2022-04-24_12-19-42.jpg")}
	c.Send("Длительность: 2,5 недели\nУроки будут высылаться через день.\nПосле каждого урока будет домашнее задание.\n\nФормат:\nВидеуроки и аудиосопровождение в закрытом Т-канале.\n\nПоддержка куратора\nПроверка домашнего задания\nНачисление баллов за нихn\n\n+ групповые созвоны с Жулдыз после каждых двух уроков\n\nИ лучшие результаты получат дополнительный приз 🎁")
	return c.Send(p, menuReserve)
}

func ReserveHandler(c tele.Context) error {
	return c.Send("Чтобы забронировать место вы можете сделать предоплату 3000 тенге переводом по номеру +77767469409\n\nПожалуйста отправьте чек об оплате 🙌🏼", mainMenu)
}

func ConnecHandler(c tele.Context) error {
	master := setMaster()
	botInstance.Send(&master, "-----")
	return c.Send("ok")
}

func WriteMaster(User tele.User) {
	jsonData, err := json.MarshalIndent(User, "", " ")
	if err != nil {
		// fuck
		log.Println("cannot serialize authHandler | 42", err)
		return
	}
	err = ioutil.WriteFile("./master.json", jsonData, 0)
	if err != nil {
		log.Println("cannot write authHandler | 47", err)
		return
	}
}

func setMaster() tele.User {
	file, err := ioutil.ReadFile("./master.json")
	if err != nil {
		log.Println("authHandler | 108", err)
	}
	master := tele.User{}
	err = json.Unmarshal(file, &master)
	if err != nil {
		log.Println("authHandler | 70", err)
	}
	fmt.Println(master)
	return master
}
