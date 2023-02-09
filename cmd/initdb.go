package main

import (
	"github.com/blogxp/ginapi/cmd/initdata"
	"github.com/blogxp/ginapi/core"
	"github.com/blogxp/ginapi/global"
	"github.com/blogxp/ginapi/initialize"

	cmd "github.com/18211167516/go-cmd"
)

// versionCmd represents the version command
var initdbCmd = &cmd.Command{
	Use:   "initdb",
	Short: "初始化表",
	Long:  `初始化核心表以及数据`,
	Run: func(Command *cmd.Command, args []string) {
		global.VP = core.Viper("../static/config/app.toml") //初始化配置
		global.VP.Set("mysql.global.LogMode", "Warn")
		global.LOG = initialize.Zap("gga") //初始化日志
		global.DB = initialize.Gorm()      //初始化DB
		initdata.InitMysqlTables(global.DB)
		initdata.InitMysqlData(global.DB)
	},
}

func init() {
	cmd.RootCmd.AddCommand(initdbCmd)

}
