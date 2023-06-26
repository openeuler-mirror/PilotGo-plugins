package global

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"gorm.io/gorm"
)

var (
	GlobalClient *client.Client
	GlobalDB     *gorm.DB
)
