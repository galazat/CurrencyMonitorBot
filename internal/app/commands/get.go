package commands

import (
	"fmt"
	"log"

	"github.com/galazat/go-telegram-bot/internal/service/currency"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	outputMsgText := "-"

	log.Println("TESTING\n\n\n\n\n")
	curr := currency.TodayCurrensies.Get(string(args))
	if curr != nil {
		outputMsgText = "sorry, can't find currency \xF0\x9F\x98\xA5"
		return
	} else {
		outputMsgText = fmt.Sprintf("\n    %d %s  -  %s RUB ", curr.Nominal, curr.CharCode, curr.Value)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsgText,
	)

	c.bot.Send(msg)
}
