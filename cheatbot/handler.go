package cheatbot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func start(ctx telebot.Context) error {
	fmt.Println(ctx.Chat().Username)
	return nil
}
