package modules

import (
	"library-management/structs"
)

func NewLibrary(id int, name string) structs.Library {
	return structs.Library{Id: id, Name: name, Users: make([]structs.User, 0), Books: make([]structs.Book, 0)}
}
