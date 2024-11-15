package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"openeuler.org/PilotGo/configmanage-plugin/db"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type ConfigInfo = internal.ConfigInfo
type ConfigNode = internal.ConfigNode
type ConfigBatch = internal.ConfigBatch
type ConfigDepart = internal.ConfigDepart

func Init() error {
	err := db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.ConfigInfo{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.ConfigNode{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.ConfigDepart{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.ConfigBatch{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.RepoFile{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.HostFile{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.SSHFile{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.SSHDFile{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.SysctlFile{})
	if err != nil {
		return err
	}
	err = db.MySQL().Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&internal.DNSFile{})
	if err != nil {
		return err
	}
	return nil
}

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
	Collect() ([]NodeResult, error)
	// 依据agent uuid进行配置下发
	Apply() ([]NodeResult, error)
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

	// 删除配置和机器关联信息，重新添加，删除的时候忽略记录不存在错误
	err = internal.DelConfigNodeByUUID(ci.UUID)
	if err != nil {
		return err
	}
	err = internal.DelConfigBatchByUUID(ci.UUID)
	if err != nil {
		return err
	}
	err = internal.DelConfigDepartByUUID(ci.UUID)
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

// 分页获取configinfo
func GetInfos(offset, size int) (int, []*ConfigInfo, error) {
	return internal.GetInfos(offset, size)
}

type Deploy struct {
	DeployBatch    common.Batch `json:"deploybatch"`
	DeployPath     string       `json:"deploypath"`
	DeployFileName string       `json:"deployname"`
	DeployText     string       `json:"deployfile"`
}

// 构造单机采集和配置下发结果结构
type NodeResult struct {
	Type     string `json:"type"`     // 指明配置类型
	NodeUUID string `json:"nodeuuid"` // 指明执行机器的uuid
	Detail   string `json:"detail"`   // 执行详情
	Result   bool   `json:"result"`   // 执行结果
	Err      string `json:"err"`      // 错误信息
}
