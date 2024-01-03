package service

import (
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

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

	// 依据agent uuid进行配置下发
	Apply(Deploy) ([]string, error)
}

type Deploy struct {
	Deploy_BatchIds  []int    `json:"deploy_batches"`
	Deploy_DepartIds []int    `json:"deploy_departs"`
	Deploy_NodeUUIds []string `json:"deploy_nodes"`
	Deploy_Path      string   `json:"deploy_path"`
	Deploy_FileName  string   `json:"deploy_name"`
	Deploy_Text      string   `json:"deploy_file"`
}

type ConfigInfo = internal.ConfigInfo
type ConfigFile = internal.ConfigFile
type ConfigNode = internal.ConfigNode
type ConfigBatch = internal.ConfigBatch
type ConfigDepart = internal.ConfigDepart

func (ci *ConfigInstance) Add(configuuid string) error {
	cm := ConfigInfo{
		UUID:           ci.UUID,
		ConfigFileUUID: configuuid,
		Type:           ci.Type,
		Description:    ci.Description,
	}
	err := cm.Add()
	if err != nil {
		return err
	}

	cb := ConfigBatch{
		ConfigInfoUUID: ci.UUID,
		BatchIDs:       ci.BatchIds,
	}
	err = cb.Add()
	if err != nil {
		return err
	}

	cd := ConfigDepart{
		ConfigInfoUUID: ci.UUID,
		DepartIDs:      ci.DepartIds,
	}
	err = cd.Add()
	if err != nil {
		return err
	}

	cn := ConfigNode{
		ConfigInfoUUID: ci.UUID,
		NodeId:         ci.Nodes,
	}
	err = cn.Add()
	if err != nil {
		return err
	}
	return nil
}

func GetInfoByConfigUUID(configuuid string) (ConfigInfo, error) {
	return internal.GetInfoByConfigUUID(configuuid)
}
