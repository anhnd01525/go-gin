package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"online-shop/config"
	"online-shop/controller"
	"online-shop/model"
	"online-shop/router"
	"os"
)

func main() {
	r := gin.Default()

	// Đọc cấu hình từ file json
	config := config.SetupConfig()


	// Khởi tạo controller
	c := controller.NewController()
	c.Config = config

	//r := setupMiddleware(ginMode)
	r.Use(cors.Default())
	
	// Kết nối CSDL
	dbConfig := config.Database
	db := model.ConnectDb(dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Address)
	defer db.Close()
	c.DB = db
	log.Println(config.Database.Database)
	err := model.MigrationDb(db, config.ServiceName)
	if err != nil {
		panic(err)
	}

	router.SetupRouter(config, r, c)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9081"
	}

	err = r.Run(":" + port)
	if err != nil {
		panic(err)
	}

}
