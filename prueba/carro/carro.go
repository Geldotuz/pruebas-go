package carro

type Auto struct {
	Name  string
	Model string
	Price float64
	Color string
	Sold  bool
}

func NewAuto(name string, model string, price float64, color string, sold bool) *Auto {
	return &Auto{
		Name:  name,
		Model: model,
		Price: price,
		Color: color,
		Sold:  sold,
	}
}

func (c *Auto) Colored(color string) {
	c.Color = color
}
