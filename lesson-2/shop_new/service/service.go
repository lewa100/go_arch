package service

import (
	"errors"
	"gb_go_arch/lesson-2/shop_new/models"
	"gb_go_arch/lesson-2/shop_new/notification"
	"gb_go_arch/lesson-2/shop_new/repository"
	"log"
)

type Service interface {
	CreateOrder(order *models.Order) (*models.Order, error)
}

type service struct {
	notif   notification.Notification
	rep     repository.Repository
	smtpBot *notification.BotSMTP
}

var (
	ErrItemNotExists = errors.New("item not exists")
)

func (s *service) CreateOrder(order *models.Order) (*models.Order, error) {
	for _, itemID := range order.ItemIDs {
		_, err := s.rep.GetItem(itemID)
		if err != nil && err != repository.ErrNotFound {
			return nil, err
		}
		if err == repository.ErrNotFound {
			return nil, ErrItemNotExists
		}
	}

	order, err := s.rep.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	//if err := s.smtpBot.Send(order); err != nil {
	//	log.Println(err)
	//}

	if err := s.notif.SendOrderCreated(order); err != nil {
		log.Println(err)
	}
	return order, err
}

func NewService(rep repository.Repository, notif notification.Notification, smtpBot notification.BotSMTP) Service {
	return &service{
		notif:   notif,
		rep:     rep,
		smtpBot: &smtpBot,
	}
}
