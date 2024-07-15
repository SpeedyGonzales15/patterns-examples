package prototype

import "encoding/json"

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	Field1 string
	Field2 int
}

func (p *ConcretePrototype) Clone() Prototype {
	clone := *p
	return &clone
}

// Глубокое копирование с использованием сериализации
func (p *ConcretePrototype) DeepCopy() Prototype {
	clone := ConcretePrototype{}
	data, _ := json.Marshal(p)
	json.Unmarshal(data, &clone)
	return &clone
}
