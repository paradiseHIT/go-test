package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var pwd_bytes []byte

type loginRequest struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

func logout(c echo.Context) error {
	sess, _ := session.Get("sd_session", c)
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

func register(c echo.Context) error {
	req := new(loginRequest)
	err := c.Bind(req)
	if err != nil {
		log.Println("Bind error", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	log.Printf("req:%v\n", req)
	// h := sha256.New()
	// h.Write([]byte(req.Password))
	// bs := h.Sum(nil)
	pwd_bytes2, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	pwd_bytes = pwd_bytes2
	log.Println(string(pwd_bytes))
	sess, _ := session.Get("sd_session", c)
	sess.Values["username"] = req.Username
	sess.Values["is_login"] = true
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/index")
}

func login(c echo.Context) error {
	log.Println("formvalue=" + c.FormValue("username"))
	req := new(loginRequest)
	err := c.Bind(req)
	if err != nil {
		log.Println("Bind error", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	log.Printf("req:%v\n", req)
	log.Print(req.Username)
	log.Print(req.Password)

	err = bcrypt.CompareHashAndPassword(pwd_bytes, []byte(req.Password))
	if err == nil {
		sess, _ := session.Get("sd_session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["username"] = req.Username
		sess.Values["is_login"] = true
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusSeeOther, "/")
	} else {
		log.Printf("err : +%v", err)
		return c.String(http.StatusOK, "user_name and password not correct")
	}
}

func index(c echo.Context) error {
	sess, _ := session.Get("sd_session", c)
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
	e.POST("/login", login)
	e.GET("/logout", logout)
	e.POST("/register", register)
	e.Logger.Fatal(e.Start("localhost:8081"))
}
