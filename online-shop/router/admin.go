package router

import (
	"github.com/gin-gonic/gin"
	"online-shop/controller"
	_ "online-shop/docs"
)

func setupAdminRoutes(c *controller.Controller, api *gin.RouterGroup) {
	api.GET("/get-users", c.GetUsers)
	api.POST("/create-user", c.CreateUser)
	api.POST("/login-user", c.UserLogIn)

	api.GET("/get-products", c.GetProducts)
	api.POST("/create-product", c.CreateProduct)
	api.GET("/get-product/:id", c.GetProductById)
	api.PUT("/update-product/:id", c.UpdateProductById)
	api.DELETE("/delete-product/:id", c.DeleteProductById)

	api.POST("/create-cart/", c.CreateCart)
	api.GET("/get-purchase-history/:id", c.GetPurchaseHistory)
}
