package main

import (
	"fmt"
	router "mvc/router"
	"log"
)

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
	fmt.Fprint(c.Writer, "hello!\n")
}

func main() {
	r := router.New()
	// 添加全局处理函数
	r.Use(GlobalHandle)

	r.GET("/", Index)

	// 增加路由分组
	r.Group("/api", func() {
		r.GET("/hello", Hello)
	}, GroupHandle)

	log.Fatal(r.Run(":8080"))
}