package modules

import (
	"library-management/structs"
	"library-management/types"
)

func NewBook(id int, name, author string, price int) types.IBook {
	return &structs.Book{Id: id, Name: name, Author: author, Price: price}
}
