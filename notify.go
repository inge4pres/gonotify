// gonotify project
package main

import (
	"fmt"
	"github.com/go-martini/martini"
	api "gonotify/gnapi"
	back "gonotify/gnbackend"
	"net/http"
	"strconv"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Recovery())

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
	m.Post("/login", func(req *http.Request, w http.ResponseWriter) {
		u, err := back.GetUserByName(req.FormValue("username"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
		}
		if u.VerifyPwd(req.FormValue("password")) {
			http.Get("/user/" + u.Uname)
		}
		w.WriteHeader(http.StatusOK)
	})

	m.Get("/user/:name", func(p martini.Params, w http.ResponseWriter) {
		user := back.GetUserByName(p["name"])
		if user.IsLogged {
			items, err := back.GetUserItems()
			for u := range items {
				w.Write([]byte(int(items[u].Id)))
			}
		} else {
			w.Write([]byte("USER not logged in"))
		}
	})

	fmt.Println("Serving on localhost:4488")
	m.RunOnAddr(":4488")
}
