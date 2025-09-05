package service

import (
	"fmt"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

type AppContext struct {
	HttpAddr string
	MySQL    *gorm.DB
	Redis    Redis
	Etcd     *clientv3.Client
	Client   *client.Client
}

var App = &AppContext{}

type AppService interface {
	Name() string
	Init(ctx *AppContext) error
	Close() error
}

type ServiceManager struct {
	services []AppService
}

func NewServiceManager(svcs ...AppService) *ServiceManager {
	return &ServiceManager{services: svcs}
}

func (sm *ServiceManager) InitAll() error {
	for _, svc := range sm.services {
		if err := svc.Init(App); err != nil {
			return fmt.Errorf("failed to init %s: %w", svc.Name(), err)
		}
		logger.Debug("Service %s initialized successfully", svc.Name())
	}
	return nil
}

func (sm *ServiceManager) CloseAll() {
	for _, svc := range sm.services {
		_ = svc.Close()
		logger.Debug("Service %s closed", svc.Name())
	}
}
