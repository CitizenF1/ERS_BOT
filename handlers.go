package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func setHandlers() {
	botInstance.Handle("/start", startHandler)
	botInstance.Handle("/users", getAllUsersFromDB)
	botInstance.Handle("/mess", sendNewsletterForUsers)

	botInstance.Handle(&btnDetail, detailHandler)
	botInstance.Handle(&btnSch, reservHandler)

	botInstance.Handle(tele.OnCallback, callBacksHandler)
	botInstance.Handle(tele.OnContact, sendContactHandler)
}

func callBacksHandler(c tele.Context) error {
	data := strings.TrimSpace(c.Data())
	identifiers := strings.Split(data, "|")
	unique := identifiers[0]
	switch unique {
	case CallbackManager1:
		botInstance.Send(c.Sender(), "Ваша заявка принята и направлена в бэк офис")
		return c.Delete()
	case CallbackManager2:
		botInstance.Send(c.Sender(), "Ваша заявка принята и направлена в бэк офис")
		return c.Delete()
	case CallbackManager3:
		botInstance.Send(c.Sender(), "Ваша заявка принята и направлена в бэк офис")
		return c.Delete()
	}
	return c.Accept()
}

func startHandler(c tele.Context) error {
	// WriteMaster(*c.Sender())
	storeUsersIntoDB(*c.Sender())
	return c.Send("Вас приветствует WANT.INVEST!\nУ нас для вас прекрасная новость: мы запускаем четвёртый поток марафона «Криптотрейдинг с нуля».\nСтоимость: 50 000 тенге\nСтарт: 29 апреля\n\nЕсли вы сегодня сделаете предоплату, цена для вас составит 38 900 тенге 🎁\n\nВам интересно наше предложение?", mainMenu)
}

func detailHandler(c tele.Context) error {
	p := &tele.Photo{File: tele.FromDisk("./img/photo_2022-04-24_12-19-42.jpg"), Caption: "Длительность: 2,5 недели\nУроки будут высылаться через день.\nПосле каждого урока будет домашнее задание.\n\nФормат:\nВидеуроки и аудиосопровождение в закрытом Т-канале.\n\nПоддержка куратора\nПроверка домашнего задания\nНачисление баллов за нихn\n\n+ групповые созвоны с Жулдыз после каждых двух уроков\n\nИ лучшие результаты получат дополнительный приз 🎁"}
	// c.Send("Длительность: 2,5 недели\nУроки будут высылаться через день.\nПосле каждого урока будет домашнее задание.\n\nФормат:\nВидеуроки и аудиосопровождение в закрытом Т-канале.\n\nПоддержка куратора\nПроверка домашнего задания\nНачисление баллов за нихn\n\n+ групповые созвоны с Жулдыз после каждых двух уроков\n\nИ лучшие результаты получат дополнительный приз 🎁")
	return c.Send(p, menuReserve)
}

func reservHandler(c tele.Context) error {
	admin := getAdminUser()
	// botInstance.Send(c.Sender(), "Чтобы забронировать место вы можете сделать предоплату 3000 тенге переводом по номеру +77767469409\n\nПожалуйста отправьте чек об оплате 🙌🏼")
	botInstance.Handle(tele.OnDocument, func(m tele.Context) error {
		botInstance.Forward(&admin, m.Message())
		return c.Send("Meнеджер с вами свяжется")
	})
	botInstance.Handle(tele.OnPhoto, func(m tele.Context) error {
		botInstance.Forward(&admin, c.Message())
		botInstance.Send(&admin, m.Message().Photo)
		return c.Send("Meнеджер с вами свяжется")
	})
	return c.Send("Чтобы забронировать место вы можете сделать предоплату 3000 тенге переводом по номеру +77767469409\n\nПожалуйста отправьте чек об оплате 🙌🏼", mainMenu)
}

func sendNewsletterForUsers(c tele.Context) error {
	admin := getAdminUser()
	// message := ""
	if c.Sender().ID == admin.ID {
		botInstance.Handle(tele.OnText, func(m tele.Context) error {
			users := usersFromStoredDB()
			for _, user := range users {
				botInstance.Send(user.User, m.Text())
			}
			return c.Send("Рассылка отправленна пользователям")
		})
	} else {
		return c.Send("Аccess denied")
	}
	return c.Send("Отправьте сообщения для рассылки")
}

func getAllUsersFromDB(c tele.Context) error {
	admin := getAdminUser()
	message := ""
	if c.Sender().ID == admin.ID {
		users := usersFromStoredDB()
		for i := 0; i < len(users); i++ {
			message += users[i].User.FirstName + " " + users[i].User.LastName + " @" + users[i].User.Username + "\n"
		}
	} else {
		botInstance.Send(c.Sender(), "Аccess denied")
	}
	return c.Send("Список пользователей: \n" + message)
}

func sendContactHandler(c tele.Context) error {
	admin := getAdminUser()
	err := c.ForwardTo(&admin)
	if err != nil {
		fmt.Println(err)
	}
	return c.Send("Meнеджер с вами свяжется")
}

func writeAdminUser(User tele.User) {
	jsonData, err := json.MarshalIndent(User, "", " ")
	if err != nil {
		log.Println("cannot serialize authHandler | 42", err)
		return
	}
	err = ioutil.WriteFile("./master.json", jsonData, 0)
	if err != nil {
		log.Println("cannot write authHandler | 47", err)
		return
	}
}

func getAdminUser() tele.User {
	file, err := ioutil.ReadFile("./master.json")
	if err != nil {
		log.Println("authHandler | 108", err)
	}
	master := tele.User{}
	err = json.Unmarshal(file, &master)
	if err != nil {
		log.Println("authHandler | 70", err)
	}
	return master
}
