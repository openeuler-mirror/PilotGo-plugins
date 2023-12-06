package service

import (
	"strconv"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type ConfigInstance struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	BatchIds    []uint   `json:"batchids"`
	DepartIds   []int    `json:"departids"`
	UUIDS       []string `json:"uuids"`
	UUID        string

	Config Config
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

type ConfigResult struct {
	ConfigInfo
	BatchIds  []uint
	DepartIds []int
	UUIDS     []string
}

func (ci *ConfigInstance) Record() error {
	cm := ConfigInfo{
		UUID:        ci.UUID,
		Type:        ci.Type,
		Description: ci.Description,
	}
	err := cm.Add()
	if err != nil {
		return err
	}

	for _, v := range ci.BatchIds {
		cn := ConfigNode{
			ConfigInfoUUID: ci.UUID,
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
			ConfigInfoUUID: ci.UUID,
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
			ConfigInfoUUID: ci.UUID,
			NodeId:         "n" + v,
		}
		err := cn.Add()
		if err != nil {
			logger.Error("save config-node failed: %s", err.Error())
			continue
		}
	}
	return err
}
