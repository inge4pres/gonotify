package gnfrontend

import back "gonotify/gnbackend"

type WebBase struct {
	Title, TitleExt string
	Status          int
	Err             error
	User            *back.User
	Items           []back.Item
}

func NewWebBase() *WebBase {
	return &WebBase{
		Title:  "GoNotify",
		Status: 200,
		User:   back.NewUser(),
	}
}
