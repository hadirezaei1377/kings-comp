package telegram

import (
	"kings-comp/internal/entity"
	"time"

	"gopkg.in/telebot.v3"
)

var (
	DefaultInputTimeout = time.Minute * 5
	DefaultTimeoutText  = `لطفا هر وقت برگشتی دوباره پیام بده.`

	TxtConfirm = `✅ بله`
	TxtDecline = `✖ خیر`
)

func GetAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}

var (
	selector           = &telebot.ReplyMarkup{}
	btnEditDisplayName = selector.Data("📝 ویرایش نام‌نمایشی", "btnEditDisplayName")
)
