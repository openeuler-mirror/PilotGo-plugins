package service

import (
	"encoding/json"

	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type ConfigInfo = internal.ConfigInfo
type ConfigNode = internal.ConfigNode
type ConfigBatch = internal.ConfigBatch
type ConfigDepart = internal.ConfigDepart

type ConfigInstance struct {
	UUID        string   `json:"uuid"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	BatchIds    []int    `json:"batchids"`
	DepartIds   []int    `json:"departids"`
	Nodes       []string `json:"uuids"`

	Config Config
}

type Config interface {
	//Version() string

	// 配置存储
	Record() error
	// 配置加载
	Load() error
	// 单机采集数据
	Collection() error
	// 依据agent uuid进行配置下发
	Apply() (json.RawMessage, error)
}

func (ci *ConfigInstance) Add() error {
	cm := ConfigInfo{
		UUID:        ci.UUID,
		Type:        ci.Type,
		Description: ci.Description,
	}
	err := cm.Add()
	if err != nil {
		return err
	}

	for _, v := range ci.Nodes {
		cn := ConfigNode{
			ConfigInfoUUID: ci.UUID,
			NodeId:         v,
		}
		err = cn.Add()
		if err != nil {
			return err
		}
	}

	for _, v := range ci.BatchIds {
		cb := ConfigBatch{
			ConfigInfoUUID: ci.UUID,
			BatchID:        v,
		}
		err = cb.Add()
		if err != nil {
			return err
		}
	}

	for _, v := range ci.DepartIds {
		cd := ConfigDepart{
			ConfigInfoUUID: ci.UUID,
			DepartID:       v,
		}
		err = cd.Add()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetInfoByUUID(configuuid string) (ConfigInfo, error) {
	return internal.GetInfoByUUID(configuuid)
}

func GetConfigByUUID(configuuid string) (*ConfigInstance, error) {
	ci, err := GetInfoByUUID(configuuid)
	if err != nil {
		return nil, err
	}

	nodes, err := internal.GetConfigNodesByUUID(configuuid)
	if err != nil {
		return nil, err
	}
	batchids, err := internal.GetConfigBatchByUUID(configuuid)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(configuuid)

	config := &ConfigInstance{
		UUID:        ci.UUID,
		Type:        ci.Type,
		Description: ci.Description,
		Nodes:       nodes,
		BatchIds:    batchids,
		DepartIds:   departids,
	}
	return config, err
}

type Deploy struct {
	Deploy_BatchIds  []int    `json:"deploy_batches"`
	Deploy_DepartIds []int    `json:"deploy_departs"`
	Deploy_NodeUUIds []string `json:"deploy_nodes"`
	Deploy_Path      string   `json:"deploy_path"`
	Deploy_FileName  string   `json:"deploy_name"`
	Deploy_Text      string   `json:"deploy_file"`
}
