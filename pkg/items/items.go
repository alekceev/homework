package items

import (
	"fmt"

	"homework/pkg/globals"
)

type Item struct {
	Id          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
}

type Items []Item

func GetAll() (Items, error) {
	rows, err := globals.Db.Query("select * from items")
	if err != nil {
		return Items{}, err
	}

	items := Items{}

	for rows.Next() {
		i := Item{}
		err := rows.Scan(&i.Id, &i.Name, &i.Description, &i.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, i)
	}

	return items, nil
}

func CreateItem(item Item) (Item, error) {
	res, err := globals.Db.Exec("insert into items(name, description, price) values (?, ?, ?)", item.Name, item.Description, item.Price)
	if err != nil {
		return item, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return item, err
	}

	item.Id = id

	return item, nil
}

func GetItem(id int64) (Item, error) {

	row := globals.Db.QueryRow("select * from items where id = ?", id)
	item := Item{}
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.Price)
	if err != nil {
		return item, err
	}
	return item, nil
}

func DeleteItem(id int64) error {
	_, err := globals.Db.Exec("delete from items where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
