package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/lie-flat-planet/srv-kubelie/api/router/demo"
	"github.com/lie-flat-planet/srv-kubelie/cmd/srv-kubelie/docs"
	"github.com/lie-flat-planet/srv-kubelie/config"
)

func NewRoot(r *gin.Engine) {
	basePath := r.Group("/srv-kubelie/api")
	v1 := basePath.Group("/v1")

	demo.Router(v1)

}

func Start() {
	r := &router{}
	r.init().registerHandler().swagger()

	config.Config.Server.GinServe(r.getEngin())
}

type router struct {
	ginEngine *gin.Engine
}

func (r *router) init() *router {
	gin.SetMode(config.Config.Server.RunMode)
	r.ginEngine = gin.Default()

	return r
}

func (r *router) registerHandler() *router {
	NewRoot(r.ginEngine)

	return r
}

func (r *router) swagger() *router {
	docs.SwaggerInfo.BasePath = "/srv-kubelie/api/v1"
	r.ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func (r *router) getEngin() *gin.Engine {
	return r.ginEngine
}