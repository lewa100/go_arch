package service

import (
	"shop/models"
	rep "shop/repository"
	"testing"
)

type dbMock struct {
	rep.RepositoryMock

	items map[int32]*models.Item
}

func (m *dbMock) CreateItem(item *models.Item) (*models.Item, error) {
	return &models.Item{
		ID:    item.ID,
		Price: item.Price,
		Name:  item.Name,
	}, nil
}

func TestCreateItem(t *testing.T) {
	s := &service{
		db: dbMock{},
	}

	item1 := &models.Item{
		Name:  "",
		Price: 342.0,
	}
	_, err := s.CreateItem(item1)
	if err.Error() != "item name is empty" {
		t.Error("expected empty error")
		return
	}

	item2 := &models.Item{
		Name:  "NotEmpty",
		Price: 0,
	}
	_, err = s.CreateItem(item2)
	if err.Error() != "item price should be positive" {
		t.Error("expected price positive error")
		return
	}
}
