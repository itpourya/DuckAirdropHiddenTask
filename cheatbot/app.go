package cheatbot

import (
	"log"

	"gopkg.in/telebot.v3"
)

func GenarateAPP() *telebot.Bot {
	setings := telebot.Settings{
		Token:  TOKEN,
		Poller: POLLER,
	}

	app, err := telebot.NewBot(setings)
	if err != nil {
		log.Fatal(err)
	}

	app.Handle("/start", start)
	app.Handle(telebot.OnText, text)

	return app
}
