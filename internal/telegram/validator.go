package telegram

import (
	"slices"

	"gopkg.in/telebot.v3"
)

func choiceValidator(choices ...string) Validator {
	return Validator{
		Validator: func(msg *telebot.Message) bool {
			return slices.Contains(choices, msg.Text)
		},
		OnInvalid: func(msg *telebot.Message) string {
			return `یکی از گزینه‌ها را انتخاب کنید`
		},
	}
}
