package logger

import (
	"os"

	"gitee.com/openeuler/PilotGo-plugin-template/conf"
	"gitee.com/openeuler/PilotGo/sdk/logger"
)

func InitLogger() {
	err := logger.Init(conf.Global_Config.Logopts)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
