package structs

import "strings"

type Library struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	// Relations
	Users   []User   `json:"users"`
	Books   []Book   `json:"books"`
	Borrows []Borrow `json:"borrows"`
}

func (l *Library) GetName() string {
	return l.Name
}

func (l *Library) GetUsers() []User {
	return l.Users
}

func (l *Library) GetBooks() []Book {
	return l.Books
}

func (l *Library) JoinUser(user *User) {
	l.Users = append(l.Users, *user)
}

func (l *Library) JoinBook(book *Book) {
	l.Books = append(l.Books, *book)
}

func (l *Library) JoinBorrow(b *Borrow) {
	l.Borrows = append(l.Borrows, *b)
}

func (l *Library) SearchBookByTitle(t string) *Book {
	for _, book := range l.Books {
		if title := strings.ToLower(book.Name); title == strings.ToLower(t) {
			return &book
		}
	}

	return nil
}
