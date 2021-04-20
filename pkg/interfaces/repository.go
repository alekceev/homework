package interfaces

import "homework/pkg/models"

type Repository interface {
	GetAll() (models.Items, error)
	Get(id int64) (*models.Item, error)
	Save(item *models.Item) error
	Delete(id int64) error
}
