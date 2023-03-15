package models

type Todo struct {
	Id        int64
	Item      string
	Completed bool
	Focused   bool
}
