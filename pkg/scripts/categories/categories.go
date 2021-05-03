package main

import (
	"fmt"
	"homework/pkg/models"
	"strings"
)

type Cats struct {
	nameToCat map[string]*models.Category
	cats      []*models.Category
}

func (c *Cats) AddCategory(id int, name string, slug string, parentName string) {
	cat := models.NewCategory(id, name, slug)

	parent, ok := c.nameToCat[parentName]
	if ok {
		parent.Append(cat)
	}
	c.cats = append(c.cats, cat)
	c.nameToCat[name] = cat
}

func main() {

	// Список категорий для примера
	c := &Cats{
		nameToCat: make(map[string]*models.Category),
	}
	c.AddCategory(1, "Главная", "", "")
	c.AddCategory(2, "Процессоры", "", "Главная")
	c.AddCategory(3, "Мат.Платы", "", "Главная")
	c.AddCategory(4, "Память", "", "Главная")
	c.AddCategory(5, "Intel", "", "Процессоры")
	c.AddCategory(6, "Amd", "", "Процессоры")
	c.AddCategory(7, "DDR3", "", "Память")
	c.AddCategory(8, "DDR4", "", "Память")
	c.AddCategory(9, "Core I3", "", "Intel")
	c.AddCategory(10, "Core I5", "", "Intel")

	cat, _ := c.nameToCat["Core I5"]

	fmt.Printf("\n\nItem: %s\nParent: %s(%d)\n", cat.Name, cat.GetParent().Name, cat.GetParent().Id)
	fmt.Printf("Root: %s\n", cat.GetRoot().Name)
	fmt.Println("Breadcrumbs: " + strings.Join(cat.Bredcrumbs().Names(), "/"))
}
