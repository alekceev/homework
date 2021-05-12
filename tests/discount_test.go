package tests

import (
	"context"
	"homework/pkg/database"
	"homework/pkg/repositories"
	"homework/pkg/services/discounter"
	"testing"
)

// не понятно, как это тестировать

func Discounter_Test(t *testing.T) {

	db, err := database.Connect("db/url")
	if err != nil {
		panic(err)
	}

	itemsRepo := repositories.NewItemRepository(db)
	service := discounter.NewDiscountService(itemsRepo, "./web/discount.csv")

	service.Start(context.Background())
}
