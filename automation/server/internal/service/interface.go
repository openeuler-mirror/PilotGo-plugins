package service

import "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"

type AppService interface {
	Name() string
	Init(ctx *global.AppContext) error
	Close() error
}
