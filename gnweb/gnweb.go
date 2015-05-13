package gnweb

import back "gonotify/gnbackend"

type BaseWeb struct {
	Title string
	Err   error
	User  *back.User
	Items []back.Item
}
