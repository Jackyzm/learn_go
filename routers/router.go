package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	// docs is generated by Swag CLI, you have to import it.
	"learn_go/docs"
	"learn_go/pkg/setting"
	"learn_go/pkg/validate"
	"learn_go/routers/api"
	"learn_go/routers/user"
)

// GetRouters all routers
func GetRouters() *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "API documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", setting.ServerSetting.HTTPPort)
	docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.HandleMethodNotAllowed = true

	router.POST("/register", user.Register)
	router.POST("/login", user.Login)

	route := router.Group("/api")
	{
		route.GET("/ping", api.GetPing)
	}

	// url := ginSwagger.URL("http://localhost:8888/swagger/doc.json") // The url pointing to API definition
	if setting.ServerSetting.RunMode == "debug" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	validate.SetUp()

	return router
}
