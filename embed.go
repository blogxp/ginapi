package main

import (
	"embed"

	"github.com/blogxp/ginapi/asset"
	"github.com/blogxp/ginapi/global"
)

//go:embed config
var rbacf embed.FS

func initEmbed() {
	global.RBACFS = rbacf                   //初始化rbac策略，只能在main中才行？
	global.STAICFS = asset.InitStaicEmbed() //初始化staic，公开
	global.TMPLFS = asset.InitTmplEmbed()   //初始化模板，不公开
}
