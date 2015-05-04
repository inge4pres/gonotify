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
		intid, _ := strconv.Atoi(p["id"])
		resp, err := api.GetItem(int64(intid))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

	m.Post("/api", func(req *http.Request, w http.ResponseWriter) {
		resp, err := api.PostItem(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

	m.Delete("/api/:id", func(p martini.Params, w http.ResponseWriter) {
		intid, _ := strconv.Atoi(p["id"])
		resp, err := api.DeleteItem(int64(intid))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(resp)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

	fmt.Println("Serving on localhost:4488")
	m.RunOnAddr(":4488")
}
