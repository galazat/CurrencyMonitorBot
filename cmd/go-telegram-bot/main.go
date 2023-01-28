package main

import (
	"log"
	"os"

	"github.com/galazat/go-telegram-bot/internal/app/commands"
	"github.com/galazat/go-telegram-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	commader := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message != nil { // If we got a message

			switch update.Message.Command() {
			case "help":
				commader.Help(update.Message)
			case "list":
				commader.List(update.Message)
			default:
				commader.Default(update.Message)
			}

		}
	}
}
