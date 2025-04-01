package main

import (
	"fmt"
	"library-management/modules"
	"library-management/structs"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var rw sync.RWMutex
	bookChan := make(chan *structs.Book)

	library1 := modules.NewLibrary(1, "library1")
	user1 := modules.NewUser(1, "Ali", "ali@gmail.com")
	book1 := modules.NewBook(1, "Who can", "A MZ", 20)
	borrow1 := modules.NewBorrow(user1, book1)

	library1.JoinBook(&book1)
	library1.JoinUser(&user1)
	library1.JoinBorrow(&borrow1)
	borrow1.ChangeStatus(true)

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer rw.RUnlock()

		rw.RLock()

		book := library1.SearchBookByTitle("who can")

		bookChan <- book
	}()

	book := <-bookChan
	if book == nil {
		err := fmt.Errorf("book is not found")
		log.Println(err)
		return
	}

	fmt.Printf("$%d\n", book.GetPrice())

	close(bookChan)
	wg.Wait()
}
