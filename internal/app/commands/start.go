package commands

import (
	"fmt"
	"time"

	"github.com/galazat/go-telegram-bot/internal/service/currency"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var START_CURRENCIES = []string{"USD", "EUR", "GBP", "UAH"}

func (c *Commander) Start(inputMessage *tgbotapi.Message) {
	tm := time.Now()

	text := fmt.Sprintf("Привет, @%s ! \xF0\x9F\x91\x8B \n\n", inputMessage.Chat.UserName)
	text += fmt.Sprintf("Данный бот показывает актуальный курс валют \xF0\x9F\x92\xB5, взятый с сайта ЦБ \xF0\x9F\x8F\xA6, по отношению к RUB на текущую дату: %s \n", tm.Format("02.01.2006"))

	for i := 0; i < len(currency.TodayCurrensies.Carrencyes); i++ {
		if isContain(currency.TodayCurrensies.Carrencyes[i].CharCode, START_CURRENCIES) {
			switch currency.TodayCurrensies.Carrencyes[i].CharCode {
			case "GBP":
				text += fmt.Sprintf("\n    1 %s \xF0\x9F\x87\xAC\xF0\x9F\x87\xA7  -  %s RUB \xF0\x9F\x87\xB7\xF0\x9F\x87\xBA", currency.TodayCurrensies.Carrencyes[i].CharCode, currency.TodayCurrensies.Carrencyes[i].Value)
			case "USD":
				text += fmt.Sprintf("\n    1 %s \xF0\x9F\x87\xBA\xF0\x9F\x87\xB8  -  %s RUB \xF0\x9F\x87\xB7\xF0\x9F\x87\xBA", currency.TodayCurrensies.Carrencyes[i].CharCode, currency.TodayCurrensies.Carrencyes[i].Value)
			case "EUR":
				text += fmt.Sprintf("\n    1 %s \xF0\x9F\x87\xA9\xF0\x9F\x87\xAA  -  %s RUB \xF0\x9F\x87\xB7\xF0\x9F\x87\xBA", currency.TodayCurrensies.Carrencyes[i].CharCode, currency.TodayCurrensies.Carrencyes[i].Value)
			case "UAH":
				text += fmt.Sprintf("\n    1 %s \xF0\x9F\x87\xA8\xF0\x9F\x87\xB3  -  %s RUB \xF0\x9F\x87\xB7\xF0\x9F\x87\xBA", currency.TodayCurrensies.Carrencyes[i].CharCode, currency.TodayCurrensies.Carrencyes[i].Value)

			}
		}
	}

	text += fmt.Sprintf("\n\nДля просмотра возможностей бота введите команду /help.\n")

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		text,
	)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Список команд", "help"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Полный список курсов валют", "list"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Конвертировать валюту", "convert"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Курс определённой валюты", "get"),
		),
	)

	c.bot.Send(msg)
}

func isContain(s string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if s == arr[i] {
			return true
		}
	}
	return false
}
