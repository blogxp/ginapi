package tool

import "github.com/blogxp/ginapi/global"

func Config(key string) interface{} {
	return global.VP.Get(key)
}
