package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/service"
)

type PaginationQ struct {
	Ok             bool        `json:"ok"`
	Size           int         `form:"size" json:"size"`
	CurrentPageNum int         `form:"page" json:"page"`
	Data           interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	TotalPage      int         `json:"total"`
}

func ConfigTypeListHandler(c *gin.Context) {
	result := []string{global.Repo, global.Host, global.SSH, global.SSHD, global.Sysctl}
	response.Success(c, result, "get config type success")
}

func ConfigHandler(c *gin.Context) {
	p := &PaginationQ{}
	// 将查询参数绑定到分页查询对象 p 中
	err := c.ShouldBindQuery(p)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	num := p.Size * (p.CurrentPageNum - 1)
	total, data, err := service.GetInfos(num, p.Size)
	if err != nil {
		response.Fail(c, "parameter error", err.Error())
		return
	}
	p.Data = data
	p.TotalPage = total

	response.Success(c, p, "get config success")
}
