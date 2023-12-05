package service

import (
	"strconv"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type ConfigInstance struct {
	Type      string   `json:"type"`
	BatchIds  []uint   `json:"batchids"`
	DepartIds []int    `json:"departids"`
	UUIDS     []string `json:"uuids"`
	Config    *Config
}

type ConfigMessage = internal.ConfigMessage
type ConfigFile = internal.ConfigFile
type ConfigNode = internal.ConfigNode

func (ci *ConfigInstance) AddConfigType() error {
	cm := ConfigMessage{
		UUID:        uuid.New().String(),
		Type:        ci.Type,
		Description: "",
	}
	err := cm.AddConfigMessage()
	if err != nil {
		return err
	}
	for _, v := range ci.BatchIds {
		cn := ConfigNode{
			ConfigMessageUUID: cm.UUID,
			NodeId:            "b" + strconv.Itoa(int(v)),
		}
		err := cn.AddConfigNode()
		if err != nil {
			logger.Error("save config-batch failed: %s", err.Error())
			continue
		}
	}
	for _, v := range ci.DepartIds {
		cn := ConfigNode{
			ConfigMessageUUID: cm.UUID,
			NodeId:            "d" + strconv.Itoa(v),
		}
		err := cn.AddConfigNode()
		if err != nil {
			logger.Error("save config-depart failed: %s", err.Error())
			continue
		}
	}
	for _, v := range ci.UUIDS {
		cn := ConfigNode{
			ConfigMessageUUID: cm.UUID,
			NodeId:            "n" + v,
		}
		err := cn.AddConfigNode()
		if err != nil {
			logger.Error("save config-node failed: %s", err.Error())
			continue
		}
	}
	cf := ConfigFile{
		ConfigMessageUUID: cm.UUID,
		Name:              "",
		File:              "",
	}
	err = cf.AddConfigFile()
	return err
}
