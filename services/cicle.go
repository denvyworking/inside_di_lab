package services

// ЦИКЛИЧЕСКИЕ ЗАВИСИМОСТИ
type ServiceA struct {
	B *ServiceB
}

type ServiceB struct {
	A *ServiceA
}

func NewServiceA(b *ServiceB) *ServiceA {
	return &ServiceA{B: b}
}

func NewServiceB(a *ServiceA) *ServiceB {
	return &ServiceB{A: a}
}
