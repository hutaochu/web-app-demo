package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/web-app-demo/myserver/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// my-server godoc
// @title my-server
// @version v0.0.1
// @BasePath /apis/v1
// @Description  This is api docs of my-server
func Run(r *gin.Engine) {
	registerSwaggerRoutes(r)
	baseGroup := r.Group("/apis/v1")

	baseGroup.Use(
		middlewares.Trace(),
		middlewares.Auth(),
		middlewares.Logger(),
		middlewares.Recover(),
	)

	registerRoutes(baseGroup)
}

func registerRoutes(r *gin.RouterGroup) {
	r.GET("/user/hello", Hello)
}

func registerSwaggerRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
