package main

import (
	"flag"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	config, err := NewConfig("config.yml")
	if err != nil {
		panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(config.TelegramAPI.APIToken)
	if err != nil {
		panic(err)
	}

	flag.Parse()
	args := flag.Args()
	for _, el := range args {
		log.Println(el)
	}
	var msg string
	switch args[0] {
	case "checkit":
		msg = checkIt(args[1], args[2])
	case "ban":
		msg = banAlert(args[1], args[2])
	}
	bot.Send(tgbotapi.NewMessage(config.TelegramAPI.ChatId, msg))

}
