package telegram

import (
	"game-walker/src"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"strings"
)

const (
	WebhookURL = "WebhookURL"
	Token      = "Token telegram bot"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Fatal(err)
	}
	//bot.Debug = true

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go func() {
		if err := http.ListenAndServe(": YOUR PORT", nil); err != nil {
			log.Fatal(err)
		}
	}()

	for update := range updates {
		go func(update tgbotapi.Update) {
			var answer tgbotapi.MessageConfig
			message := update.Message.Text
			log.Printf("пришло сообщение от %v: %v ", update.Message.From.UserName, message)
			words := strings.Fields(message)

			switch words[0] {
			case "/start":
				answer = commandStart(&update)
			case "/help":
				answer = commandHelp(&update)
			case "ожить":
				_, ok := src.Players[update.Message.From.UserName]
				if !ok {
					answer = commandCreatePlayer(&update)
					player := src.AadPlayer(update.Message.From.UserName, update.Message.Chat.ID)
					go func(p *src.Player) {
						output := p.GetOutput()
						for text := range output {
							log.Printf("отпавка сообщения игроку %v: %v \n", p.Name, text)
							bot.Send(tgbotapi.NewMessage(p.ChatID, text))
						}
					}(player)
				} else {
					answer = commandPlayerBe(&update)
				}
			case "осмотреться", "идти", "одеть", "взять", "применить", "сказать", "сказать_игроку":
				player, ok := src.Players[update.Message.From.UserName]
				if ok {
					player.HandleInput(message)
				} else {
					answer = commandPlayerHaveNotLife(&update)
				}
			default:
				answer = commandUnknown(&update)
			}
			if answer.Text != "" {
				log.Printf("отпавка сообщения игроку %v: %v \n", update.Message.From.UserName, answer.Text)
				bot.Send(answer)
			}

		}(update)
	}
}

func commandPlayerBe(update *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, msgPlayerBe)
}

func commandPlayerHaveNotLife(update *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, msgPlayerHaveNotLife)
}

func commandStart(update *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, msgHello)
}

func commandUnknown(update *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, msgUnknown)
}

func commandHelp(update *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, msgHelp)
}

func commandCreatePlayer(update *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, msgCreatePlayer)
}
