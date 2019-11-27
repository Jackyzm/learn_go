package routers

import (

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

    "learn_go/docs" // docs is generated by Swag CLI, you have to import it.
    "learn_go/routers/api"
)

// GetRouters all routers
func GetRouters() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "API documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8888"
	docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

    route := router.Group("/api")
    {
        route.GET("/ping", api.GetPing)
    }

	url := ginSwagger.URL("http://localhost:8888/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":8888") // listen and serve on 0.0.0.0:8888 (for windows "localhost:8080")
}
