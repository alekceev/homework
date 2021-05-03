package categories

import "fmt"

type Category struct {
	Id        int
	Name      string
	Slug      string
	ParentId  int
	parent    *Category
	childrens []*Category
	level     int
}

func NewCategory(id int, name string, slug string) *Category {
	return &Category{
		Id:        id,
		Name:      name,
		Slug:      slug,
		childrens: make([]*Category, 0),
		level:     1,
	}
}

// Добавляем категорю к родителю, а так же назначаем категории родителя
func (c *Category) Append(category *Category) {
	c.childrens = append(c.childrens, category)
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
