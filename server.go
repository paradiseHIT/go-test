package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var cnt int

type request struct {
	Id string `json:"id" form:"id" query:"id"`
}
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func getUser(c echo.Context) error {
	// User ID from path `/users?id=123`
	//log.Println("query:" + c.QueryParam("id"))
	req := new(request)
	err := c.Bind(req)
	if err != nil {
		log.Println("Bind error", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	cnt++
	log.Printf("%d:%v\n", cnt, req)
	//go BackgroundProcess()
	return c.JSON(http.StatusOK, req)
}
func BackgroundProcess() (int, error) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		time.Sleep(time.Second)
		log.Println("Sleep ", i)
	}
	return sum, nil
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func main() {
	e := echo.New()
	e.HideBanner = true

	//e.POST("/users", saveUser)
	e.GET("/show", show)
	e.GET("/users", getUser)
	e.POST("/users", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users2", func(c echo.Context) error {
		u := new(User)
		//调用echo.Context的Bind函数将请求参数和User对象进行绑定。
		if err := c.Bind(u); err != nil {
			return err
		}
		//请求参数绑定成功后 u 对象就保存了请求参数。
		//这里直接将请求参数以json格式显示
		//注意：User结构体,字段标签定义中，json定义的字段名，就是User对象转换成json格式对应的字段名。
		return c.JSON(http.StatusCreated, u)
	})
	e.Logger.Fatal(e.Start("localhost:8080"))
}
