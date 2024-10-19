package cheatbot

import (
	"time"

	"gopkg.in/telebot.v3"
)

var (
	TOKEN  = "TOKEN"
	POLLER = &telebot.LongPoller{Timeout: 1 * time.Second}
)
