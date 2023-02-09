package global

import (
	"embed"
	// "github.com/blogxp/ginapi/config"

	cron "github.com/18211167516/robfig-cron/v3"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CF      Config
	VP      *viper.Viper
	DB      *gorm.DB
	LOG     *zap.Logger
	CRONLOG *zap.SugaredLogger
	CRON    *cron.Cron
	SER     Server
	Verify  *validator.Validate
	STAICFS embed.FS
	TMPLFS  embed.FS
	RBACFS  embed.FS
)

type Server interface {
	ListenAndServe() error
	Restart() error
	Shutdown()
}
