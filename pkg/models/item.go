package models

type Item struct {
	Id          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Number      string  `json:"number" db:"number"`
	Category    string  `json:"category" db:"category"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	SalePrice   float64 `json:"sale_price" db:"sale_price"`
	Amount      int64   `json:"amount" db:"amount"`
}

type Items []*Item

func (i *Item) AmountText() string {
	switch {
	case i.Amount <= 0:
		return "Нет в наличии"
	case i.Amount > 10:
		return "Много"
	default:
		return "Мало"
	}
}
