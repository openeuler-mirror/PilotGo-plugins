package service

import (
	"fmt"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
)

type ServiceManager struct {
	services []AppService
}

func NewServiceManager(svcs ...AppService) *ServiceManager {
	return &ServiceManager{services: svcs}
}

func (sm *ServiceManager) InitAll() error {
	for _, svc := range sm.services {
		if err := svc.Init(global.App); err != nil {
			return fmt.Errorf("failed to init %s: %w", svc.Name(), err)
		}
		logger.Debug("Service %s initialized successfully", svc.Name())
	}
	return nil
}

func (sm *ServiceManager) CloseAll() {
	for _, svc := range sm.services {
		_ = svc.Close()
	}
}
