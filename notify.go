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
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/js", "./static/js")

	r.GET("/api/:id", apiGet)
	r.POST("/api/", apiPost)
	r.PUT("/api/:id", apiPut)
	r.DELETE("/api/:id", apiDelete)

	r.POST("/signup", postSignup)
	r.GET("/login", getLogin)
	r.POST("/login", postLogin)

	r.GET("/user/:name", getUser)

	r.GET("/", getIndex)

	fmt.Println("Serving on localhost:4488")
	http.ListenAndServe(":4488", r)
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func getUser(c *gin.Context) {
	user, err := back.GetUserByName(c.Params.ByName("name"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", err)
	}
	if items, err := back.GetUserItems(user); err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", err)
	} else {
		c.HTML(http.StatusOK, "base.tmpl", items)
	}
}
func getLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "LOGGED IN")
	c.HTML(http.StatusOK, "login.tmpl", nil)
}
func postLogin(c *gin.Context) {
	u, err := back.GetUserByName(c.Request.FormValue("username"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.HTMLString(http.StatusInternalServerError, err.Error())
	}
	if u.VerifyPwd(c.Request.FormValue("password")) {
		c.JSON(http.StatusOK, "VERIFIED")
		c.Redirect(http.StatusMovedPermanently, "/user/"+u.Uname)
	} else {
		c.JSON(http.StatusUnauthorized, "NOT VERIFIED")
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
func postSignup(c *gin.Context) {
	uname := c.Request.FormValue("username")
	rname := c.Request.FormValue("realname")
	mail := c.Request.FormValue("mail")
	pwd := c.Request.FormValue("password")
	if err := back.RegisterUser(uname, rname, mail, pwd); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.HTML(http.StatusInternalServerError, "error.tmpl", err)

	}
	c.JSON(http.StatusAccepted, "OK")
	c.HTML(http.StatusAccepted, "index.tmpl", nil)
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
