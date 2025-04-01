package types

type IBorrow interface {
	changeStatus(status bool)
}
