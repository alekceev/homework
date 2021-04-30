package main

import (
	"fmt"
	"homework/pkg/categories"
	"strings"
)

func main() {

	categories := categories.NewCategories()
	categories.AddCategory(1, "Главная", "", "")
	categories.AddCategory(2, "Процессоры", "", "Главная")
	categories.AddCategory(3, "Мат.Платы", "", "Главная")
	categories.AddCategory(4, "Память", "", "Главная")
	categories.AddCategory(5, "Intel", "", "Процессоры")
	categories.AddCategory(6, "Amd", "", "Процессоры")
	categories.AddCategory(7, "DDR3", "", "Память")
	categories.AddCategory(8, "DDR4", "", "Память")
	categories.AddCategory(9, "Core I3", "", "Intel")
	categories.AddCategory(10, "Core I5", "", "Intel")

	cat := categories.GetCategory(10)
	fmt.Printf("\n\nItem: %s\nParent: %s(%d)\n", cat.Name, cat.GetParent().Name, cat.GetParent().Id)
	fmt.Printf("Root: %s\n", cat.GetRoot().Name)
	fmt.Println("Breadcrumbs: " + strings.Join(cat.Bredcrumbs(), "/"))

	fmt.Println("\n\nMenu")
	categories.PrintTree(1)
	fmt.Println("\n\nSubmenu")
	categories.PrintTree(2)

}
