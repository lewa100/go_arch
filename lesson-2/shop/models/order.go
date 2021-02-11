package models

type Order struct {
	ID      int32   `json:"id"`
	Phone   string  `json:"phone"`
	Email	string	`json:"email"`
	ItemIDs []int32 `json:"item_ids"`
}
