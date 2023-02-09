package initialize

import (
	"github.com/blogxp/ginapi/routes"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	return routes.InitRouter()
}
