package service

import (
	"gitee.com/openeuler/PilotGo/sdk/logger"
)

type LoggerService struct {
	Conf *logger.LogOpts
}

func (m *LoggerService) Name() string {
	return "Logger"
}
func (m *LoggerService) Init(ctx *AppContext) error {
	if err := logger.Init(m.Conf); err != nil {
		return err
	}
	return nil
}

func (m *LoggerService) Close() error {
	return nil
}
