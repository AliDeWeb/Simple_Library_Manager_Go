package main

import (
	"fmt"
	"library-management/modules"
)

func main() {
	user1 := modules.NewUser(1, "Ali", "ali@gmail.com")

	fmt.Print(user1.GetEmail())
}
