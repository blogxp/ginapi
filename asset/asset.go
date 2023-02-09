package asset

import (
	"embed"
)

//go:embed static
var staticf embed.FS

//go:embed templates
var tmplf embed.FS

func InitStaicEmbed() embed.FS {
	return staticf //初始化FS
}
func InitTmplEmbed() embed.FS {
	return tmplf //初始化FS
}
