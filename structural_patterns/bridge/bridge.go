package bridge

// Implementor определяет интерфейс для классов реализации.
type Implementor interface {
	OperationImpl() string
}

// Abstraction определяет абстрактный интерфейс и содержит ссылку на объект Implementor.
type Abstraction struct {
	imp Implementor
}

func (a *Abstraction) Operation() string {
	return a.imp.OperationImpl()
}

// ConcreteImplementorA и ConcreteImplementorB представляют конкретные реализации, которые следуют интерфейсу Implementor.
type ConcreteImplementorA struct{}

func (c *ConcreteImplementorA) OperationImpl() string {
	return "ConcreteImplementorA OperationImpl"
}

type ConcreteImplementorB struct{}

func (c *ConcreteImplementorB) OperationImpl() string {
	return "ConcreteImplementorB OperationImpl"
}
