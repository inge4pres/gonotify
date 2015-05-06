// gonotify project
package main

import (
	"fmt"
	"github.com/go-martini/martini"
	api "gonotify/gnapi"
	"net/http"
	"strconv"
)

func main() {
	m := martini.Classic()

	m.Get("/api/:id", func(p martini.Params, w http.ResponseWriter) {
		resp, _ := api.GetItem(p["id"])
		w.Write(resp)
	})
	m.Post("/api", func(req *http.Request, w http.ResponseWriter) {
		resp, _ := api.PostItem(req)
		w.Write(resp)
	})
	m.Delete("/api/:id", func(p martini.Params, w http.ResponseWriter) {
		intid, _ := strconv.Atoi(p["id"])
		resp, _ := api.DeleteItem(int64(intid))
		w.Write(resp)
	})

	fmt.Println("Serving on localhost:4488")
	m.RunOnAddr(":4488")
}
