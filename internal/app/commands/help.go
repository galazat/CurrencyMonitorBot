package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/start - стартовое приветствие. Показаны курсы основных валют\n"+
			"/help - показывает список доступных команд\n\n"+
			"/list - показывает полный список курсов валют\n\n"+
			"/get - показывает курс определённой валюты. Необходимо ввести код валюты, пример:  /get USD \n\n"+
			"/convert - конвентирует сумму в RUB в сумму выбранной валюты. Пример запроса: /convert 1000->USD\nКонвенрирует 1000 RUB в 14.15 USD (при курсе 1 USD = 70.65 RUB) ",
	)

	c.bot.Send(msg)
}
