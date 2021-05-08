package repositories

import (
	"fmt"
	"homework/pkg/interfaces"
	"homework/pkg/models"
)

type ItemRepository struct {
	Db interfaces.DB
}

// todo***
// ItemRepositoryRaw => SQL. .Raw().Query("SELECT...")
// ItemRepository: Find() => repoRaw.Find()
// ItemRepository: Save() => repoRaw.Update() / repoRaw.Insert()
// SOFT_DELETE: created_at, updated_at... deleted_at
// ItemRepository: Delete() => repoRaw.Update(..deleted_at=...) / repoRaw.Delete()

var _ interfaces.Repository = &ItemRepository{}

func NewItemRepository(db interfaces.DB) *ItemRepository {
	return &ItemRepository{Db: db}
}

func (i *ItemRepository) GetAll() (models.Items, error) {
	rows, err := i.Db.Raw().Query("select * from items")
	if err != nil {
		return nil, err
	}

	items := models.Items{}

	for rows.Next() {
		i := &models.Item{}
		err := rows.Scan(&i.Id, &i.Name, &i.Description, &i.Number, &i.Category, &i.Price, &i.SalePrice)

		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, i)
	}

	return items, nil
}

func (i *ItemRepository) Get(id int64) (*models.Item, error) {
	row := i.Db.Raw().QueryRow("select * from items where id = ?", id)
	item := &models.Item{}
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.Number, &item.Category, &item.Price, &item.SalePrice)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (i *ItemRepository) Save(item *models.Item) error {
	salePrice := item.SalePrice
	if salePrice <= 0 {
		salePrice = item.Price
	}
	res, err := i.Db.Raw().Exec("insert into items(name, description, number, category, price, sale_price) values (?, ?, ?, ?, ?, ?)",
		item.Name, item.Description, item.Number, item.Category, item.Price, salePrice)
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
	_, err := i.Db.Raw().Exec("delete from items where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (i *ItemRepository) FindByName(name string) (*models.Item, error) {
	row := i.Db.Raw().QueryRow("select * from items where name = ?", name)
	found_item := &models.Item{}
	err := row.Scan(&found_item.Id, &found_item.Name, &found_item.Description, &found_item.Number, &found_item.Category, &found_item.Price, &found_item.SalePrice)
	if err != nil {
		return found_item, err
	}
	return found_item, nil
}

func (i *ItemRepository) Update(item *models.Item) error {
	found_item, err := i.Get(item.Id)
	if err != nil || found_item == nil {
		return fmt.Errorf("not found item %d", item.Id)
	}

	salePrice := item.SalePrice
	if salePrice <= 0 {
		salePrice = item.Price
	}

	_, err = i.Db.Raw().Exec("update items set name = ?, description = ?, number = ?, category = ?, price = ?, sale_price = ? where id = ?",
		item.Name, item.Description, item.Number, item.Category, item.Price, salePrice, item.Id)
	if err != nil {
		return err
	}

	return nil
}
