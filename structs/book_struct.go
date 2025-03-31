package structs

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

func (book *Book) GetName() string {
	return book.Name
}

func (book *Book) GetAuthor() string {
	return book.Author
}

func (book *Book) GetPrice() int {
	return book.Price
}
