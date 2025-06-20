package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDB()
	r := gin.Default()

	r.POST("/books", CreateBook)
	r.GET("/books", GetBooks)
	r.GET("/books/:id", GetBookByID)
	r.DELETE("/books/:id", DeleteBook)

	r.Run(":8080") // http://localhost:8080
}
