package models

type Supplier struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewSupplier(name string) *Supplier {
	return &Supplier{Name: name}
}

