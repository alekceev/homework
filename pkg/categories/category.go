package categories

import "fmt"

type Category struct {
	Id        int
	Name      string
	Slug      string
	ParentId  int
	parent    *Category
	childrens []*Category
	hasChild  map[int]struct{}
	level     int
}

func NewCategory(id int, name string, slug string) *Category {
	return &Category{
		Id:        id,
		Name:      name,
		Slug:      slug,
		childrens: make([]*Category, 0),
		// hasChild   make(map[int]struct{}, 0), // Не понял, почему ругается
		level: 1,
	}
}

// Добавляем категорю к родителю, а так же назначаем категории родителя
func (c *Category) Append(category *Category) {
	c.childrens = append(c.childrens, category)
	if c.hasChild == nil {
		c.hasChild = make(map[int]struct{}, 0)
	}
	c.hasChild[category.Id] = struct{}{}
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

func (c *Category) Bredcrumbs() []string {

	breadcrumbs := make([]string, 0)
	if c.parent != nil {
		breadcrumbs = append(breadcrumbs, c.parent.Bredcrumbs()...)
	}

	breadcrumbs = append(breadcrumbs, c.Name)

	return breadcrumbs
}
