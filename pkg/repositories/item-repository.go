package repositories

import (
	"database/sql"
	"fmt"
	"homework/pkg/interfaces"
	"homework/pkg/models"
)

type ItemRepository struct {
	db *sql.DB
}

var _ interfaces.Repository = &ItemRepository{}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (i *ItemRepository) GetAll() (models.Items, error) {
	rows, err := i.db.Query("select * from items")
	if err != nil {
		return nil, err
	}

	items := models.Items{}

	for rows.Next() {
		i := &models.Item{}
		err := rows.Scan(&i.Id, &i.Name, &i.Description, &i.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, i)
	}

	return items, nil
}

func (i *ItemRepository) Get(id int64) (*models.Item, error) {
	row := i.db.QueryRow("select * from items where id = ?", id)
	item := &models.Item{}
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.Price)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (i *ItemRepository) Save(item *models.Item) error {
	res, err := i.db.Exec("insert into items(name, description, price) values (?, ?, ?)", item.Name, item.Description, item.Price)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	item.Id = id

	return nil
}

func (i *ItemRepository) Delete(id int64) error {
	_, err := i.db.Exec("delete from items where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
