package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"encoding/json"
	"os"
	"strings"
//	"fmt"
	"flag"

	"finder/md"
)


var (
	config = flag.String("config", "/etc/tg-markdown-finder.json", "path to config file")
)

type Config struct {
	TelegramBotToken string
	PastePath map[string] string
}

func main() {

	flag.Parse()

	file, _ := os.Open(*config)
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 45

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		command := update.Message.Command()

		if command == "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сам ты " + update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

		} else {

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			args := strings.Fields(update.Message.Text)
			firstArg := strings.Join(args[1:], " ")

			targetFile := configuration.PastePath[command]
			markdown, _ := md.ParseMarkdownFile(targetFile)

			switch firstArg {
			case "list":

				anchors := md.GetAllAnchors(markdown)
				anchors_str := strings.Join(anchors, " ")
				
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, anchors_str)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			default: 
				paste := firstArg

				content := md.GetBlockByAnchors(markdown, paste)
				content_str := strings.Join(content, " ")

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, content_str)
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ParseMode = "Markdown"

				bot.Send(msg)
			}
		}
	}
}