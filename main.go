package main

import (
	"library-management/modules"
)

func main() {
	library1 := modules.NewLibrary(1, "library1")
	user1 := modules.NewUser(1, "Ali", "ali@gmail.com")
	book1 := modules.NewBook(1, "Who can pass the road?", "A MZ", 20)
	borrow1 := modules.NewBorrow(user1, book1)

	library1.JoinBook(&book1)
	library1.JoinUser(&user1)
	library1.JoinBorrow(&borrow1)
	borrow1.ChangeStatus(true)

}
