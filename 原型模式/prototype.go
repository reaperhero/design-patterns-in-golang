package Prototype

type Prototype interface {
	GetName() string
	SetName(name string)
	Clone() Prototype
}

type ConcretePrototype struct {
	name string
}

func (p *ConcretePrototype) GetName() string {
	return p.name
}

func (p *ConcretePrototype) SetName(name string) {
	p.name = name
}

func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{name: p.name}
}
