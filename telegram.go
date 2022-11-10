package main

import (
	"log"

	tb "gopkg.in/telebot.v3"
)

func tgSendMessage(msg string, chatID int) (responce *tb.Message) {
	tbot, err := tb.NewBot(tb.Settings{
		Token: cfg.Telegram.BotToken,
	})
	if err != nil {
		log.Println(err)
	} else {
		group := tb.ChatID(chatID)
		var opts tb.SendOptions
		opts.ParseMode = tb.ModeHTML
		responce, err = tbot.Send(group, msg, &opts)
		if err != nil {
			log.Println(err)
			log.Println(msg)
		}
	}
	return responce
}
