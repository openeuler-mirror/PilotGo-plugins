package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigNode struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	NodeId         string `json:"node_id"` //机器uuid
}

func (cn *ConfigNode) Add() error {
	return db.MySQL().Save(&cn).Error
}

func GetConfigNodesByUUID(uuid string) ([]string, error) {
	var nodes []string
	err := db.MySQL().Where("config_info_uuid=?", uuid).Select("node_id").Find(&nodes).Error
	return nodes, err
}

type ConfigBatch struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	BatchID        int `json:"batch_id"`
}

func (cb *ConfigBatch) Add() error {
	return db.MySQL().Save(&cb).Error
}

func GetConfigBatchByUUID(uuid string) ([]int, error) {
	var nodes []int
	err := db.MySQL().Where("config_info_uuid=?", uuid).Select("batch_id").Find(&nodes).Error
	return nodes, err
}

type ConfigDepart struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	DepartID       int `json:"depart_id"`
}

func (cd *ConfigDepart) Add() error {
	return db.MySQL().Save(&cd).Error
}

func GetConfigDepartByUUID(uuid string) ([]int, error) {
	var nodes []int
	err := db.MySQL().Where("config_info_uuid=?", uuid).Select("depart_id").Find(&nodes).Error
	return nodes, err
}
