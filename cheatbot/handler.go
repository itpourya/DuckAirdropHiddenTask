package cheatbot

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v3"
)

func start(ctx telebot.Context) error {
	fmt.Println(ctx.Chat().Username)
	return nil
}

func text(ctx telebot.Context) error {
	context := ctx.Text()

	if !strings.HasPrefix(context, "Bearer ") {
		return ctx.Send("bad token")
	}

	return ctx.Send("Accept Token")
}
