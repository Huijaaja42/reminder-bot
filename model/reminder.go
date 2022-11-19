package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Reminder struct {
	Id      uint64
	User    string
	Channel string
	Time    int64
	Text    string
}
