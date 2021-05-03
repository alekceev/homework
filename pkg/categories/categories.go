package categories

type Categories struct {
	Cats []*Category
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
