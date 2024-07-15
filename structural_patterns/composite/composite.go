package composite

// Component определяет интерфейс для всех компонентов, включая как базовые, так и составные объекты.
type Component interface {
	Operation() string
}

// Leaf представляет базовые объекты структуры, которые не имеют подкомпонентов.
type Leaf struct{}

func (l *Leaf) Operation() string {
	return "Leaf"
}

// Composite хранит группу дочерних компонентов и реализует связанные с ними операции.
type Composite struct {
	children []Component
}

func (c *Composite) Add(child Component) {
	c.children = append(c.children, child)
}

func (c *Composite) Operation() string {
	result := "Composite with children: "
	for _, child := range c.children {
		result += child.Operation() + " "
	}
	return result
}
