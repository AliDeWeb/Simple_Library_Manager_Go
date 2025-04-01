package main

import (
	"fmt"
	"library-management/modules"
)

func main() {
	library1 := modules.NewLibrary(1, "library1")

	user1 := modules.NewUser(1, "Ali", "ali@gmail.com")
	fmt.Println(user1.GetEmail())

	book1 := modules.NewBook(1, "Who can pass the road?", "A MZ", 20)
	fmt.Println(book1.GetPrice())

	library1.JoinBook(&book1)
	library1.JoinUser(&user1)

	fmt.Println(library1.GetUsers())
	fmt.Println(library1.GetBooks())
}
