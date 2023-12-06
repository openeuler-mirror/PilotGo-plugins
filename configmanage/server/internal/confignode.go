package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigNode struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	NodeId         string
}

func (cn *ConfigNode) Add() error {
	return db.MySQL().Save(&cn).Error
}

func GetConfigNodesByUUID(uuid string) ([]ConfigNode, error) {
	var nodes []ConfigNode
	err := db.MySQL().Where("config_message_uuid=?", uuid).Find(&nodes).Error
	return nodes, err
}
