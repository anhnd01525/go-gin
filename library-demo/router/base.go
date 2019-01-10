package router

import (
	//"github.com/swaggo/gin-swagger"
	//"log"
	//"os"
	//"os/exec"
	"time"

	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-gonic/gin"
	"library-demo/config"
	"library-demo/controller"
)

type JsonDataRoute struct {
	Tags        []string `json:"tags"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
}

type JsonData struct {
	Paths map[string]map[string]JsonDataRoute
}

type Route struct {
	Service     string
	Endpoint    string
	HttpMethod  string
	Description string
	IsPublic    bool
	IsAdmin     bool
}

type Message struct {
	Id     string
	Routes []Route
	Time   time.Time
}

func SetupRouter(config config.Config, r *gin.Engine, c *controller.Controller) {
	// Mọi routes đều bắt đầu bởi prefix ServiceName
	api := r.Group("/" + config.ServiceName)
	{
		setupAdminRoutes(c, api)

		//setupDocumentRoute(api, "")
	}
}

//func setupDocumentRoute(api *gin.RouterGroup, ginMode string) {
//	if ginMode != "release" {
//		goPath := os.ExpandEnv("$GOPATH")
//		if goPath == "" {
//			panic("Chưa thiết lập GOPATH!")
//		}
//
//		_, err := exec.Command("/"+goPath+"/bin/swag", "init").Output()
//
//		if err != nil {
//			log.Println("Không thể tạo document bằng Swag!")
//			panic(err)
//		}
//
//		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//	}
//}
