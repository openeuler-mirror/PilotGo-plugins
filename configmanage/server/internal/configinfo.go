package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigInfo struct {
	ID          int    `gorm:"unique;autoIncrement:true"`
	UUID        string `gorm:"type:varchar(50);primary_key" json:"uuid"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (cm *ConfigInfo) Add() error {
	return db.MySQL().Create(cm).Error
}

func GetInfoByUUID(uuid string) (ConfigInfo, error) {
	var ci ConfigInfo
	err := db.MySQL().Where("uuid=?", uuid).Find(&ci).Error
	return ci, err
}
