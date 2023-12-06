package service

import (
	"strconv"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type ConfigInstance struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	BatchIds    []uint   `json:"batchids"`
	DepartIds   []int    `json:"departids"`
	UUIDS       []string `json:"uuids"`

	Config *Config
}

type Config interface {

	// 配置存储
	Record() error
	// 配置加载
	Load() error

	// 依据agent uuid进行配置下发
	Apply(string) error
}

type ConfigInfo = internal.ConfigInfo
type ConfigFile = internal.ConfigFile
type ConfigNode = internal.ConfigNode

func (ci *ConfigInstance) Record() (string, error) {
	cm := ConfigInfo{
		UUID:        uuid.New().String(),
		Type:        ci.Type,
		Description: ci.Description,
	}
	err := cm.Add()
	if err != nil {
		return cm.UUID, err
	}

	for _, v := range ci.BatchIds {
		cn := ConfigNode{
			ConfigInfoUUID: cm.UUID,
			NodeId:         "b" + strconv.Itoa(int(v)),
		}
		err := cn.Add()
		if err != nil {
			logger.Error("save config-batch failed: %s", err.Error())
			continue
		}
	}

	for _, v := range ci.DepartIds {
		cn := ConfigNode{
			ConfigInfoUUID: cm.UUID,
			NodeId:         "d" + strconv.Itoa(v),
		}
		err := cn.Add()
		if err != nil {
			logger.Error("save config-depart failed: %s", err.Error())
			continue
		}
	}

	for _, v := range ci.UUIDS {
		cn := ConfigNode{
			ConfigInfoUUID: cm.UUID,
			NodeId:         "n" + v,
		}
		err := cn.Add()
		if err != nil {
			logger.Error("save config-node failed: %s", err.Error())
			continue
		}
	}
	return cm.UUID, err
}
