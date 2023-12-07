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

func GetInfoByConfigUUID(configuuid string) (ConfigInfo, error) {
	var ci ConfigInfo
	err := db.MySQL().Where("config_file_uuid=?", configuuid).Find(&ci).Error
	return ci, err
}

func GetInfoByUUID(uuid string) ([]ConfigInfo, error) {
	var cis []ConfigInfo
	err := db.MySQL().Where("uuid=?", uuid).Find(&cis).Error
	return cis, err
}
