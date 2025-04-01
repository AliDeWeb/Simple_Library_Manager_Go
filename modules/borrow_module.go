package modules

import "library-management/structs"

func NewBorrow(user structs.User, book structs.Book) structs.Borrow {
	return structs.Borrow{User: user, Book: book, Returned: false}
}
