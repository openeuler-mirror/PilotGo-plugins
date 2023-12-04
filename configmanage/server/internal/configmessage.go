package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigMessage struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT"`
	UUID        string `json:"uuid"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (cm *ConfigMessage) AddConfigMessage() error {
	return db.MySQL().Save(&cm).Error
}

func GetConfigMessage() ([]ConfigMessage, error) {
	var cm []ConfigMessage
	err := db.MySQL().Find(&cm).Error
	return cm, err
}
