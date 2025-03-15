package services

type Prototype interface {
	Clone() Prototype
}
