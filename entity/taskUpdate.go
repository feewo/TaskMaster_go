package entity

type TaskUpdate struct {
	Tid   uint32
	Title string
	Ready bool
	Iid   uint32
}
