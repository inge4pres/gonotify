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
	m.Patch("/api/:id", func(p martini.Params, w http.ResponseWriter) {
		intid, _ := strconv.Atoi(p["id"])
		resp, _ := api.ArchiveItem(int64(intid))
		w.Write(resp)
	})
	m.Delete("/api/:id", func(p martini.Params, w http.ResponseWriter) {
		intid, _ := strconv.Atoi(p["id"])
		resp, _ := api.DeleteItem(int64(intid))
		w.Write(resp)
	})
	m.Post("/register", func(req *http.Request, w http.ResponseWriter) {
		uname := req.FormValue("username")
		rname := req.FormValue("realname")
		mail := req.FormValue("mail")
		pwd := req.FormValue("password")
		if err := back.RegisterUser(uname, rname, mail, pwd); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("REGISTERED"))
	})
	m.Post("/login", func(req *http.Request, w http.ResponseWriter) {
		u, err := back.GetUserByName(req.FormValue("username"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}
		if u.VerifyPwd(req.FormValue("password")) {
			w.WriteHeader(http.StatusOK)
			http.Get("/user/" + u.Uname)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}

	})
	m.Get("/user/:name", func(p martini.Params, w http.ResponseWriter) {
		user, err := back.GetUserByName(p["name"])
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		if user.IsLogged {
			items, _ := back.GetUserItems(user)
			for u := range items {
				w.Write(api.RenderJson(items[u]))
			}
		} else {
			w.Write([]byte("USER not logged in"))
			http.Get("/login")
		}
	})

	fmt.Println("Serving on localhost:4488")
	m.RunOnAddr(":4488")
}
