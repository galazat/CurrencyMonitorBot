package commands

import (
	"fmt"

	"github.com/galazat/go-telegram-bot/internal/service/currency"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Полный список курсов валют: \n\n"

	for i := 0; i < len(currency.TodayCurrensies.Carrencyes); i++ {
		outputMsgText += fmt.Sprintf("\n    %d %s  -  %s RUB ", currency.TodayCurrensies.Carrencyes[i].Nominal, currency.TodayCurrensies.Carrencyes[i].CharCode, currency.TodayCurrensies.Carrencyes[i].Value)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Список команд", "help"),
		),
	)

	c.bot.Send(msg)
}
