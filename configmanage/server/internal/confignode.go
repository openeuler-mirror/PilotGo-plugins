package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigNode struct {
	ID             int    `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfoUUID string `json:"config_info_uuid"`
	NodeId         string `json:"node_id"`
}

func (cn *ConfigNode) Add() error {
	return db.MySQL().Save(&cn).Error
}

func GetConfigNodesByUUID(uuid string) ([]ConfigNode, error) {
	var nodes []ConfigNode
	err := db.MySQL().Where("config_info_uuid=?", uuid).Find(&nodes).Error
	return nodes, err
}
