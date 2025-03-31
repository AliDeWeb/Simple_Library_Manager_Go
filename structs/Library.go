package structs

type Library struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	// Relations
	Users []User `json:"users"`
	Books []Book `json:"books"`
}
