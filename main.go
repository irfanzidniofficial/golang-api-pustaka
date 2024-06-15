package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {

	router:=gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)
	
	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"name": "Muhammad Irfan Zidni",
			"bio": "A Software Engineer",
		});
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello",
		"bio": "Belajar Golang",
	})
}