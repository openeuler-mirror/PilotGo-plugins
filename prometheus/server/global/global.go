package global

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"gorm.io/gorm"
)

var (
	GlobalClient        *client.Client
	GlobalDB            *gorm.DB
	GlobalPrometheusYml string
)

const (
	GlobalPrometheusYmlInit = "../scripts/init_prometheus_yml.sh"
)
