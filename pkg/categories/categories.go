package categories

import (
	"fmt"
)

type Categories struct {
	ids      map[int]*Category
	nameToId map[string]int
}

func NewCategories() *Categories {
	return &Categories{
		ids:      make(map[int]*Category),
		nameToId: make(map[string]int),
	}
}

func (c *Categories) AddCategory(id int, name string, slug string, parent_name string) {
	if _, found := c.ids[id]; found {
		fmt.Printf("Category exists %#v\n", found)
		return
	}

	category := NewCategory(id, name, slug)

	parentId, ok := c.nameToId[parent_name]
	if ok {
		parent := c.ids[parentId]
		parent.Append(category)
	}

	c.ids[category.Id] = category
	c.nameToId[category.Name] = category.Id
}

func (c *Categories) GetCategory(id int) *Category {
	return c.ids[id]
}

func (c *Categories) PrintTree(id int) {
	cat := c.GetCategory(id)
	if cat == nil {
		return
	}

	// ограничение на уровень вложенности
	if cat.level >= 5 {
		return
	}

	if len(cat.childrens) == 0 {
		return
	}

	tabs := ""
	for i := 1; i < cat.level; i++ {
		tabs += "\t"
	}

	for _, child := range cat.childrens {
		fmt.Printf("%s%s(%d)\n", tabs, child.Name, child.Id)
		c.PrintTree(child.Id)
	}
}
