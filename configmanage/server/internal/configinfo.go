package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigInfo struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	UUID           string     `gorm:"type:varchar(50)" json:"uuid"`
	ConfigFile     ConfigFile `gorm:"Foreignkey:ConfigFileUUID"`
	ConfigFileUUID string
	Type           string `json:"type"`
	Description    string `json:"description"`
}

func (cm *ConfigInfo) Add() error {
	return db.MySQL().Save(&cm).Error
}

func GetConfigMessage() ([]ConfigInfo, error) {
	var cm []ConfigInfo
	err := db.MySQL().Find(&cm).Error
	return cm, err
}
