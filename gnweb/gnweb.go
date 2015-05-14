package gnweb

import back "gonotify/gnbackend"

type Base struct {
	Title, TitleExt string
	Status          int
	Err             error
	User            *back.User
	Items           []back.Item
}

func New() *Base {
	return &Base{
		Title:  "GoNotify",
		Status: 200,
	}
}
