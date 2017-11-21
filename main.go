package main

import (
	"encoding/json"
	"fmt"
	router "go-router/router"
	"log"
)

//
type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

var st = &Student{
	"Xiao Ming",
	16,
	true,
	[]string{"Math", "English", "Chinese"},
	9.99,
}

// GlobalHandle 全局处理函数
func GlobalHandle(c *router.Context) {
	fmt.Fprint(c.Writer, "begin GlobalHandle!\n")
	c.Next()
	fmt.Fprint(c.Writer, "end GlobalHandle!\n")
}

func Index(c *router.Context) {
	fmt.Fprint(c.Writer, "Welcome!\n")
}

// GroupHandle 分组处理函数
func GroupHandle(c *router.Context) {
	fmt.Fprint(c.Writer, "begin GroupHandle!\n")
	c.Next()
	fmt.Fprint(c.Writer, "end GroupHandle!\n")
}

func Hello(c *router.Context) {
	fmt.Fprint(c.Writer, "hello1!\n")
}

func Test(c *router.Context) {
	b, err := json.Marshal(st)
	if err == nil {
		//fmt.Fprint(c.Writer, string(b))
		//data := string(b)
		c.JSON(b)
		//fmt.Fprint(c.Writer, data)
	} else {
		fmt.Fprint(c.Writer, "1111!\n")
	}
}

func main() {
	r := router.New()
	// 添加全局处理函数
	r.Use(GlobalHandle)

	r.GET("/", Index)
	// 增加路由分组
	r.Group("/api", func() {
		r.GET("/hello", Hello)
		r.GET("/test", Test)
	}, GroupHandle)
	log.Fatal(r.Run(":8088"))
}
