// gonotify project
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	api "gonotify/gnapi"
	back "gonotify/gnbackend"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/api/:id", apiGet)
	r.POST("/api/", apiPost)
	r.PUT("/api/:id", apiPut)
	r.DELETE("/api/:id", apiDelete)

	r.POST("/register", register)
	//	m.Post("/login", func(req *http.Request, w http.ResponseWriter) {
	//		u, err := back.GetUserByName(req.FormValue("username"))
	//		if err != nil {
	//			w.WriteHeader(http.StatusNotFound)
	//			w.Write([]byte(err.Error()))
	//		}
	//		if u.VerifyPwd(req.FormValue("password")) {
	//			w.WriteHeader(http.StatusOK)
	//			http.Redirect(w, req, "/user/"+u.Uname, 301)
	//		} else {
	//			w.WriteHeader(http.StatusUnauthorized)
	//		}

	//	})
	//	m.Get("/user/:name", func(p martini.Params, w http.ResponseWriter) {
	//		user, err := back.GetUserByName(p["name"])
	//		if err != nil {
	//			w.Write([]byte(err.Error()))
	//		}
	//		if user.IsLogged {
	//			items, _ := back.GetUserItems(user)
	//			w.WriteHeader(http.StatusOK)
	//			w.Write([]byte("You have " + strconv.Itoa(len(items)) + " notification(s) "))
	//			for u := range items {
	//				w.Write(api.RenderJson(items[u]))
	//			}
	//		} else {
	//			w.WriteHeader(http.StatusUnauthorized)
	//			w.Write([]byte("USER not logged in"))
	//		}
	//	})

	fmt.Println("Serving on localhost:4488")
	http.ListenAndServe(":4488", r)
}

func register(c *gin.Context) {
	uname := c.Request.FormValue("username")
	rname := c.Request.FormValue("realname")
	mail := c.Request.FormValue("mail")
	pwd := c.Request.FormValue("password")
	if err := back.RegisterUser(uname, rname, mail, pwd); err != nil {
		c.HTMLString(http.StatusInternalServerError, "Error : %s", err.Error())
	}
	c.HTMLString(http.StatusAccepted, "Registered with username %s!", uname)
}

func apiGet(c *gin.Context) {
	resp := api.GetItem(c.Params.ByName("id"))
	c.JSON(resp.Status, resp)
}
func apiPost(c *gin.Context) {
	resp := api.GetItem(c.Params.ByName("id"))
	c.JSON(resp.Status, resp)
}
func apiPut(c *gin.Context) {
	resp := api.PostItem(c.Request)
	c.JSON(resp.Status, resp)
}
func apiDelete(c *gin.Context) {
	resp := api.DeleteItem(c.Params.ByName("id"))
	c.JSON(resp.Status, resp)
}
