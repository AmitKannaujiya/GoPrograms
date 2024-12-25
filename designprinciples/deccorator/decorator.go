package deccorator

type Coffee struct {
	Price int
}

func (c *Coffee) GetPrice() int {
	return c.Price
}

type FilterCoffee struct {
	Coffee Coffee
}

func (c *FilterCoffee) GetPrice() int {
	return c.Coffee.GetPrice() * 2
}

type CreamCoffee struct {
	Coffee Coffee
}

func (c *CreamCoffee) GetPrice() int {
	return c.Coffee.GetPrice() * 2 + 5
}