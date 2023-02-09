package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/blogxp/ginapi/app/services"
	"github.com/blogxp/ginapi/tool"
)

// @Summary 首页
// @Description  首页
// @Router /admin/index [get]
func Index(c *gin.Context) {
	list := services.GetLeftMenuList(c.GetString("uid"))
	user, _ := c.Get("waitUse")
	data := tool.M{
		"title":                "黑龙江齐齐哈尔",
		"list":                 list["data"],
		"userInfo":             user,
		"changeOwnPasswordUrl": "/admin/changeOwnPassword",
		"loginoutUrl":          "/admin/loginout",
		"userView":             "/admin/userView",
	}
	tool.HTML(c, "common/index.html", data)
}

func Main(c *gin.Context) {
	tool.HTML(c, "common/main.html", nil)
}
