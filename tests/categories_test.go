package tests

import (
	"homework/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Beadcrumbs(t *testing.T) {
	item := models.NewCategory(1, "item", "item")
	subItem := models.NewCategory(2, "Subitem", "Subitem")
	item.Append(subItem)

	expect := models.NewCategories().Add(item).Add(subItem)
	assert.Equal(t, expect, subItem.Bredcrumbs())
}
