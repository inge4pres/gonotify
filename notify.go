// gonotify project
package main

import (
	"fmt"
	"github.com/go-martini/martini"
	back "gonotify/gnbackend"
	"net/http"
)

var be back.Item

func main() {
	m := martini.Classic()
	m.Post("/", func(req *http.Request) string {
		be.Rcpt_mail = req.FormValue("mail")
		be.Level = "INFO"
		err := back.RecvItem(be)
		if err != nil {
			return err.Error()
		}
		return "Added item with rcpt_mail " + be.Rcpt_mail + "\n"
	})
	fmt.Println("Serving on localhost:4488")
	m.RunOnAddr(":4488")
}
