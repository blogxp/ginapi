package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/blogxp/ginapi/app/middleware"
	"github.com/blogxp/ginapi/global"
)

//init router
func InitRouter() *gin.Engine {
	r := initGin()
	loadRoute(r)
	return r
}

// init Gin
func initGin() *gin.Engine {
	//设置gin模式
	gin.SetMode(global.VP.GetString("RunMode"))
	engine := gin.New()
	engine.Use()
	//定义404中间件
	engine.NoRoute(middleware.NoRoute())
	return engine
}

// 加载路由
func loadRoute(r *gin.Engine) {
	fileRoute(r)
	testRoute(r)
	apiRoute(r)
	swagRoute(r)
	adminRoute(r)
}
