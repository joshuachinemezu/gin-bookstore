package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joshuachinemezu/gin-bookstore/controllers"
	"github.com/joshuachinemezu/gin-bookstore/models"
)

func handleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)

	r.Run()
}
