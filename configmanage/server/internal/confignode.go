package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigNode struct {
	ID             int    `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfoUUID string `json:"config_info_uuid"`
	NodeId         string `json:"node_id"` //机器uuid
}

type ConfigBatch struct {
	ID             int    `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfoUUID string `json:"config_info_uuid"`
	BatchID        int    `json:"batch_id"`
}

type ConfigDepart struct {
	ID             int    `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfoUUID string `json:"config_info_uuid"`
	DepartID       int    `json:"depart_id"`
}

func (cn *ConfigNode) Add() error {
	return db.MySQL().Save(&cn).Error
}

func (cb *ConfigBatch) Add() error {
	return db.MySQL().Save(&cb).Error
}

func (cd *ConfigDepart) Add() error {
	return db.MySQL().Save(&cd).Error
}

func GetConfigNodesByUUID(uuid string) (ConfigNode, error) {
	var nodes ConfigNode
	err := db.MySQL().Where("config_info_uuid=?", uuid).Find(&nodes).Error
	return nodes, err
}

func GetConfigBatchByUUID(uuid string) (ConfigBatch, error) {
	var nodes ConfigBatch
	err := db.MySQL().Where("config_info_uuid=?", uuid).Find(&nodes).Error
	return nodes, err
}

func GetConfigDepartByUUID(uuid string) (ConfigDepart, error) {
	var nodes ConfigDepart
	err := db.MySQL().Where("config_info_uuid=?", uuid).Find(&nodes).Error
	return nodes, err
}
