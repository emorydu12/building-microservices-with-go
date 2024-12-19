package data

type Store interface {
	Search(name string) []Kitten
}