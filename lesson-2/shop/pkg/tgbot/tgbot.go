package tgbot

import (
	"errors"
	"fmt"
	"net/http"
	"shop/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	ErrChatNotFound = errors.New("chat not found")
)

type TelegramAPI interface {
	SendOrderNotification(order *models.Order) error
}

type telegramAPI struct {
	tgBot  *tgbotapi.BotAPI
	chatId int64
}

func NewTelegramAPI(token string, chatID int64) (*telegramAPI, error) {
	cli := &http.Client{
		Timeout: 10 * time.Second,
	}
	bot, err := tgbotapi.NewBotAPIWithClient(token, cli)
	if err != nil {
		return nil, err
	}
	return &telegramAPI{
		tgBot:  bot,
		chatId: chatID,
	}, nil
}

func (s *telegramAPI) SendOrderNotification(order *models.Order) error {
	text := fmt.Sprintf("new order %d\nemail: %s\nphone: %s", order.ID, order.Email, order.Phone)

	msg := tgbotapi.NewMessage(s.chatId, text)

	_, err := s.tgBot.Send(msg)
	return err
}
