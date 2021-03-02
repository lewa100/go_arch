package repository

import (
	"shop/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	db := NewMapDB()

	type structForTest struct {
		input  models.Item
		create models.Item
		update models.Item
	}

	dataForTest := []structForTest{
		{
			input: models.Item{
				Name:  "someName1",
				Price: 10,
			},
			create: models.Item{
				ID:    1,
				Name:  "someName1",
				Price: 10,
			},
			update: models.Item{
				ID:    1,
				Name:  "someName10",
				Price: 100,
			},
		},
		{
			input: models.Item{
				Name:  "someName2",
				Price: 20,
			},
			create: models.Item{
				ID:    2,
				Name:  "someName2",
				Price: 20,
			},
			update: models.Item{
				ID:    2,
				Name:  "someName20",
				Price: 200,
			},
		},
		{
			input: models.Item{
				Name:  "someName3",
				Price: 40,
			},
			create: models.Item{
				ID:    3,
				Name:  "someName3",
				Price: 40,
			},
			update: models.Item{
				ID:    3,
				Name:  "someName30",
				Price: 400,
			},
		},
		{
			input: models.Item{
				Name:  "someName4",
				Price: 50,
			},
			create: models.Item{
				ID:    4,
				Name:  "someName4",
				Price: 50,
			},
			update: models.Item{
				ID:    4,
				Name:  "someName40",
				Price: 500,
			},
		},
	}

	t.Run("TestCreateItem", func(t *testing.T) {
		for _, v := range dataForTest {
			result, err := db.CreateItem(&v.input)
			if err != nil {
				t.Error("unexpected error: ", err)
			}
			assert.Equal(t, v.create.ID, result.ID, "unexpected name: expected %d result: %d", result.ID, v.create.ID)
			assert.Equal(t, v.create.Name, result.Name, "unexpected name: expected %d result: %d", result.Name, v.create.Name)
			assert.Equal(t, v.create.Price, result.Price, "unexpected name: expected %d result: %d", result.Price, v.create.Price)
		}
	})

	t.Run("TestUpdateItem", func(t *testing.T) {
		for _, v := range dataForTest {
			result, err := db.UpdateItem(&v.update)
			if err != nil {
				t.Error("unexpected error: ", err)
			}
			assert.Equal(t, v.update.ID, result.ID, "unexpected name: expected %d result: %d", result.ID, v.update.ID)
			assert.Equal(t, v.update.Name, result.Name, "unexpected name: expected %d result: %d", result.Name, v.update.Name)
			assert.Equal(t, v.update.Price, result.Price, "unexpected name: expected %d result: %d", result.Price, v.update.Price)
		}
	})

	t.Run("TestDeleteItem", func(t *testing.T) {
		err := db.DeleteItem(2)
		if err != nil {
			t.Error("unexpected error: ", err)
		}
		_, err = db.GetItem(2)
		assert.EqualError(t, err, "not found", "Item not found")
	})

	type structForItemFilterTest struct {
		in  ItemFilter
		out int
	}

	listForTest := []structForItemFilterTest{
		{
			in: ItemFilter{
				Limit: 4,
			},
			out: 3,
		}, {
			in: ItemFilter{
				//PriceLeft:  createInt64(20),
				//PriceRight: createInt64(50),
				Limit: 3,
			},
			out: 3,
		},
		{
			in: ItemFilter{
				Limit: 1,
			},
			out: 1,
		},
	}

	t.Run("TestListItems", func(t *testing.T) {
		for _, v := range listForTest {
			r, err := db.ListItems(&v.in)
			if err != nil {
				t.Error("unexpected error: ", err)
			}
			assert.Equal(t, len(r), v.out, "unexpected name: expected %d result: %d", r, v)
			//for _, it := range r {
			//	assert.True(t, *v.in.PriceLeft <= it.Price, "unexpected name: The condition is not met %d <= %d <= %d", v.in.PriceLeft, it, v.in.PriceRight)
			//}
		}
	})
}

//func createInt64(x int) *int64 {
//	tmp := int64(x)
//	return &tmp
//}
