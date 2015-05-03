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

	m.Post("/api", func(req *http.Request, w http.ResponseWriter) {
		id, err := api.PostItem(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(int(id))))
	})

	m.Delete("/api/:id", func(p martini.Params, w http.ResponseWriter) {
		intid, _ := strconv.Atoi(p["id"])
		err := api.DeleteItem(int64(intid))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("Serving on localhost:4488")
	m.RunOnAddr(":4488")
}
