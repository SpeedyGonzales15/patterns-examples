package factory

type Product interface {
	Use() string
}

type ConcreteProduct struct {
}

func (p *ConcreteProduct) Use() string {
	return "ConcreteProduct"
}

type Creator interface {
	CreateProduct() Product
}

type ConcreteCreator struct{}

func (c *ConcreteCreator) CreateProduct() Product {
	return &ConcreteProduct{}
}
