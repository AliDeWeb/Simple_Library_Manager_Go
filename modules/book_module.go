package modules

import (
	"library-management/structs"
)

func NewBook(id int, name, author string, price int) structs.Book {
	return structs.Book{Id: id, Name: name, Author: author, Price: price}
}
