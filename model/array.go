package model

type Element interface {
	IsValid() bool
}

type Array interface {
	FilterEmptyEntries() Array
}
