package router

import (
	"github.com/gin-gonic/gin"
	"library-demo/controller"
)

func setupAdminRoutes(c *controller.Controller, api *gin.RouterGroup) {
	api.GET("/get-users", c.GetUsers)
	api.POST("/create-user", c.CreateUser)

	api.GET("/get-authors", c.GetAuthors)
	api.POST("/create-author", c.CreateAuthor)
	api.PUT("/update-author/:id", c.UpdateAuthorById)

	api.GET("/get-types", c.GetTypes)
	api.POST("/create-type", c.CreateType)
	api.PUT("/update-type/:id", c.UpdateTypeById)

	api.GET("/get-publishers", c.GetPublishers)
	api.POST("/create-publisher", c.CreatePublisher)
	api.PUT("/update-publisher/:id", c.UpdatePublisherById)

	api.GET("/get-books", c.GetBooks)
	api.POST("/create-book", c.CreateBook)
	api.PUT("/update-book/:id", c.	UpdateBookById)
	api.DELETE("/delete-book/:id", c.DeleteBookById)
	api.GET("/get-book/:id", c.GetBookById)

	api.POST("/borrow-books", c.BorrowBooks)
	api.POST("/return-books", c.ReturnBooks)
}
