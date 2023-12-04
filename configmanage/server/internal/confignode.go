package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigNode struct {
	ID                int           `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigMessage     ConfigMessage `gorm:"Foreignkey:ConfigMessageUUID"`
	ConfigMessageUUID string
	NodeId            string
}

func (cn *ConfigNode) AddConfigNode() error {
	return db.MySQL().Save(&cn).Error
}

func GetConfigDodesByUUID(uuid string) ([]ConfigNode, error) {
	var nodes []ConfigNode
	err := db.MySQL().Where("config_message_uuid=?", uuid).Find(&nodes).Error
	return nodes, err
}
