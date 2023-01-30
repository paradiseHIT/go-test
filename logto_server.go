package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Id string `json:"id" form:"id" query:"id"`
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	is_login := sess.Values["is_login"]
	str := fmt.Sprintf("%v", is_login)
	log.Println("sess.Values['is_login'] in logout = " + str)
	if is_login != true {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	sess.Values["username"] = ""
	sess.Values["is_login"] = ""
	sess.Save(c.Request(), c.Response())
	return c.String(http.StatusOK, "logout_page")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	if username == "u" && password == "p" {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["username"] = username
		sess.Values["is_login"] = true
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusSeeOther, "/index")
	} else {
		return c.String(http.StatusOK, "user_name and password not correct")
	}
}

func index(c echo.Context) error {
	sess, _ := session.Get("session", c)
	is_login := sess.Values["is_login"]
	str := fmt.Sprintf("%v", is_login)
	log.Println("sess.Values['is_login'] in index = " + str)
	if is_login != true {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	return c.String(http.StatusOK, "index_page")
}

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/index", index)
	e.GET("/login", login)
	e.GET("/logout", logout)
	e.Logger.Fatal(e.Start("localhost:8081"))
}
