package main

import (
	"fmt"
	"library-management/modules"
)

func main() {
	user1 := modules.NewUser(1, "Ali", "ali@gmail.com")
	fmt.Println(user1.GetEmail())

	book1 := modules.NewBook(1, "Who can pass the road?", "A MZ", 20)
	fmt.Println(book1.GetPrice())
}
