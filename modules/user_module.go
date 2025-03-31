package modules

import (
	"library-management/structs"
	"library-management/types"
)

func NewUser(id int, name, email string) types.IUser {
	return &structs.User{Id: id, Name: name, Email: email}
}
