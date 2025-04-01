package types

import "library-management/structs"

type ILibrary interface {
	GetName() string
	GetUsers() []structs.User
	GetBooks() []structs.Book

	JoinUser(user *structs.User)
	JoinBook(book *structs.Book)
	JoinBorrow(borrow *structs.Borrow)

	SearchBookByTitle(t string) *structs.Book
}
