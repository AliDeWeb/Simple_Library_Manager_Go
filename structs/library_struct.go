package structs

type Library struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	// Relations
	Users []User `json:"users"`
	Books []Book `json:"books"`
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
