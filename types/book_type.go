package types

type IBook interface {
	GetName() string
	GetAuthor() string
	GetPrice() int
}
