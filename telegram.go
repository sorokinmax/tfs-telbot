package main

import (
	"log"

	tb "gopkg.in/telebot.v3"
)

func tgSendMessage(msg string) (responce *tb.Message) {
	tbot, err := tb.NewBot(tb.Settings{
		Token: cfg.Telegram.BotToken,
	})
	if err != nil {
		log.Println(err)
	} else {
		group := tb.ChatID(cfg.Telegram.BuildChatID)
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
