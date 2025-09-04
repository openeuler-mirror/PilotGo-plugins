package app

import (
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
)

type LoggerService struct {
	Conf *logger.LogOpts
}

func (m *LoggerService) Name() string {
	return "Logger"
}
func (m *LoggerService) Init(ctx *global.AppContext) error {
	if err := logger.Init(m.Conf); err != nil {
		return err
	}
	return nil
}

func (m *LoggerService) Close() error {
	return nil
}
