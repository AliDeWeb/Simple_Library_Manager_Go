package structs

type Borrow struct {
	User     User `json:"user"`
	Book     Book `json:"book"`
	Returned bool `json:"returned"`
}

func (b *Borrow) ChangeStatus(status bool) {
	b.Returned = status
}
