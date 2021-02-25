package notification

import (
	"fmt"
	"gb_go_arch/lesson-2/shop_new/models"
	"net/smtp"
)

type BotSMTP struct {
	from    string `json:"from"`
	host    string `json:"host"`
	port    string `json:"port"`
	address string `json:"address"`
	passwd  string `json:"password"`
}

func NewSMTPBot(from, host, port, passwd string) *BotSMTP {
	return &BotSMTP{
		from,
		host,
		port,
		host + ":" + port,
		passwd,
	}
}

func (bot *BotSMTP) Send(o *models.Order) error {
	subject := fmt.Sprintf("New Order §%d", o.ID)
	body := fmt.Sprintf("new order %d\n\nphone: %s\nemail: %s", o.ID, o.CustomerPhone, o.CustomerEmail)

	message := []byte(subject + "\n" + body)
	auth := smtp.PlainAuth("", bot.from, bot.passwd, bot.host)

	if err := smtp.SendMail(bot.address, auth, bot.from, []string{o.CustomerEmail}, message); err != nil {
		return err
	}
	return nil
}
