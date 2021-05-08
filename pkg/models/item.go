package models

import (
	"reflect"
)

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

func (item *Item) DbColNames() []string {

	colNames := make([]string, 0)
	val := reflect.Indirect(reflect.ValueOf(item))

	for i := 0; i < val.Type().NumField(); i++ {

		col := val.Type().Field(i).Tag.Get("db")

		// without primary key
		if col == "id" {
			continue
		}

		colNames = append(colNames, col)
	}
	return colNames
}
