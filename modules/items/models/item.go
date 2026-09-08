package models

type Item struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SupplierId  int     `json:"supplier_id"`
}

func NewItem(name, desc string, price float32, supp int) *Item {
	return &Item{
		Name:        name,
		Description: desc,
		Price:       price,
		SupplierId:  supp,
	}
}
