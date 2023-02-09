package response

import (
	"github.com/blogxp/ginapi/app/models"
)

type SysRuleResponse struct {
	models.SysRule
	Checked bool `json:"checked" gorm:"-"`
}
