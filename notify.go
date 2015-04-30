// gonotify project
package main

import (
	"fmt"
	"github.com/go-martini/martini"
	hand "gonotify/handlers"
	"net/http"
)

var m *martini.Martini

func main() {
	m = martini.New()

	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	r := martini.NewRouter()

	r.Post("/", hand.PostItem(req*http.Request))
	fmt.Println("Serving on localhost:4488")
	m.RunOnAddr(":4488")
}
