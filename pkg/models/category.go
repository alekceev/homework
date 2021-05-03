package models

import "fmt"

type Category struct {
	Id       int
	Name     string
	Slug     string
	ParentId int
	parent   *Category
	children []*Category
	level    int
}

type Categories struct {
	Cats []*Category
}

func NewCategory(id int, name string, slug string) *Category {
	return &Category{
		Id:       id,
		Name:     name,
		Slug:     slug,
		children: make([]*Category, 0),
		level:    1,
	}
}

// Добавляем категорю к родителю, а так же назначаем категории родителя
func (c *Category) Append(category *Category) {
	c.children = append(c.children, category)
	category.ParentId = c.Id
	category.parent = c
	category.level = c.level + 1
}

func (c *Category) GetParent() *Category {
	return c.parent
}

func (c *Category) GetRoot() *Category {
	checked := make(map[int]struct{}, 0)
	for {
		if _, ok := checked[c.Id]; ok {
			panic(fmt.Sprintf("Recursion on id: %d", c.Id))
		}
		if c.parent == nil {
			return c
		}
		checked[c.Id] = struct{}{}
		c = c.parent
	}
}

func (c *Category) Bredcrumbs() *Categories {

	breadcrumbs := NewCategories()
	if c.parent != nil {
		for _, cat := range c.parent.Bredcrumbs().Cats {
			breadcrumbs.Add(cat)
		}
	}

	breadcrumbs.Add(c)

	return breadcrumbs
}

func NewCategories() *Categories {
	return &Categories{
		Cats: make([]*Category, 0),
	}
}

func (c *Categories) Add(cat *Category) *Categories {
	c.Cats = append(c.Cats, cat)
	return c
}

func (c *Categories) Names() []string {
	names := make([]string, 0)
	if len(c.Cats) == 0 {
		return names
	}
	for _, cat := range c.Cats {
		names = append(names, cat.Name)
	}
	return names
}
