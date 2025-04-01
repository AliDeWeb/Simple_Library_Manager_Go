package modules

import (
	"library-management/structs"
)

func NewUser(id int, name, email string) structs.User {
	return structs.User{Id: id, Name: name, Email: email}
}
