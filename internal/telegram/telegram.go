package telegram

import (
	"kings-comp/internal/service"
	"kings-comp/internal/telegram/teleprompt"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot

	TelePrompt *teleprompt.TelePrompt
}

func NewTelegram(app *service.App, apiKey string) (*Telegram, error) {

	t := &Telegram{
		App:        app,
		TelePrompt: teleprompt.NewTelePrompt(),
	}
	pref := telebot.Settings{
		Token:   apiKey,
		Poller:  &telebot.LongPoller{Timeout: 60 * time.Second},
		OnError: t.onError,
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Error("couldn't connect to telegram servers")
		return nil, err
	}

	t.bot = bot

	t.setupHandlers()
	return t, nil
}

func (t *Telegram) Start() {
	t.bot.Start()
}
