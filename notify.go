// gonotify project
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	back "gonotify/gnbackend"
	fe "gonotify/gnfrontend"
	se "gonotify/gnsession"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/js", "./static/js")

	r.GET("/api/:id", validSession, apiGet)
	r.POST("/api/", validSession, apiPost)
	r.PUT("/api/:id", apiPut)
	r.DELETE("/api/:id", apiDelete)

	r.GET("/signup", getSignup)
	r.POST("/signup", postSignup)
	r.GET("/login", getLogin)
	r.POST("/login", postLogin)
	r.GET("/logout", validSession, logOut)

	r.GET("/user/:name", validSession, getUser)

	r.GET("/", getIndex)

	fmt.Println("Serving on localhost:4488")
	http.ListenAndServe(":4488", r)
}

func getIndex(c *gin.Context) {
	w := fe.New()
	c.HTML(http.StatusOK, "index.tmpl", &w)
}
func getUser(c *gin.Context) {
	w := fe.New()
	user, err := back.GetUserByName(c.Params.ByName("name"))
	if err != nil {
		w.Err = err
	}
	items, err := back.GetUserItems(user)
	if err != nil {
		w.Status = http.StatusInternalServerError
		w.Err = err
	} else {
		w.Items = items
	}
	w.Title += " - " + user.Uname
	w.User = user
	c.HTML(w.Status, "base.tmpl", &w)
}
func getLogin(c *gin.Context) {
	w := fe.New()
	//	c.JSON(http.StatusOK, "LOGGED IN")
	c.HTML(http.StatusOK, "login.tmpl", &w)
}
func postLogin(c *gin.Context) {
	u, err := back.GetUserByName(c.Request.FormValue("username"))
	if err != nil {
		w := fe.New()
		w.Err = err
		w.Status = http.StatusInternalServerError
		//c.JSON(w.Status, err)
		c.HTML(w.Status, "base.tmpl", &w)
	}
	if u.VerifyPwd(c.Request.FormValue("password")) {
		setSessionCookie(c, u)
		c.Redirect(http.StatusMovedPermanently, "/user/"+u.Uname)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
func getSignup(c *gin.Context) {
	w := fe.New()
	c.HTML(http.StatusOK, "signup.tmpl", &w)
}
func postSignup(c *gin.Context) {
	w := fe.New()
	uname := c.Request.FormValue("username")
	rname := c.Request.FormValue("realname")
	mail := c.Request.FormValue("email")
	pwd := c.Request.FormValue("password")
	user, err := back.RegisterUser(uname, rname, mail, pwd)
	if err != nil {
		w.Err = err
		w.Status = http.StatusInternalServerError
	}
	w.User = user
	setSessionCookie(c, user)
	c.HTML(w.Status, "index.tmpl", &w)
}

//API
func apiGet(c *gin.Context) {
	resp := fe.GetItem(c.Params.ByName("id"))
	c.JSON(resp.Status, resp)
}
func apiPost(c *gin.Context) {
	resp := fe.PostItem(c.Request)
	c.JSON(resp.Status, resp)
}
func apiPut(c *gin.Context) {
	resp := fe.ArchiveItem(c.Params.ByName("id"))
	c.JSON(resp.Status, resp)
}
func apiDelete(c *gin.Context) {
	resp := fe.DeleteItem(c.Params.ByName("id"))
	c.JSON(resp.Status, resp)
}

//SESSION
func setSessionCookie(c *gin.Context, u *back.User) {
	session := se.New()
	if err := session.CreateSession(u); err != nil {
		w := fe.New()
		w.Err = err
		c.HTML(http.StatusInternalServerError, "base.tmpl", &w)
	}
	http.SetCookie(c.Writer, session.Scookie)
}
func validSession(c *gin.Context) {
	cookie, err := c.Request.Cookie("sessionid")
	if err != nil || cookie == nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	} else {
		if se.VerifyCookie(cookie) {
			c.Writer.WriteHeader(http.StatusContinue)
		}
	}
}
func logOut(c *gin.Context) {
	cookie, _ := c.Request.Cookie("sessionid")
	err := se.Logout(cookie)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/")
	}
	c.Redirect(http.StatusOK, "/")
}
