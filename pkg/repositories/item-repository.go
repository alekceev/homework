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
	items := models.Items{}

	err := i.Db.Raw().Select(&items, "select * from items")
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *ItemRepository) Get(id int64) (*models.Item, error) {
	item := &models.Item{}
	err := i.Db.Raw().Get(&item, "select * from items where id = ?", id)
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

	res, err := i.Db.Raw().NamedExec("insert into items(name, description, number, category, price) values (:name, :desctiption, :number, :category, :price)", &item)
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
	item := &models.Item{}
	err := i.Db.Raw().Get(&item, "select * from items where name = ?", name)
	if err != nil {
		return item, err
	}
	return item, nil
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

	_, err = i.Db.Raw().NamedExec("update items set name = :name, description = :description, number = :number, category = :category, price = :price, sale_price = :sale_price where id = :id",
		&item)
	if err != nil {
		return err
	}

	return nil
}
