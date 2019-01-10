package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"library-demo/config"
	"library-demo/controller"
	"library-demo/model"
	"library-demo/router"
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
		port = "9082"
	}

	err = r.Run(":" + port)
	if err != nil {
		panic(err)
	}

}
