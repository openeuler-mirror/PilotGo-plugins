package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigInfo struct {
	ID          int    `gorm:"AUTO_INCREMENT"`
	UUID        string `gorm:"type:varchar(50);primary_key" json:"uuid"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Info2File struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	ConfigFile     ConfigFile `gorm:"Foreignkey:ConfigFileUUID"`
	ConfigFileUUID string
	Version        string `gorm:"type:varchar(50)" json:"version"`
}

func (cm *ConfigInfo) Add() error {
	return db.MySQL().Create(cm).Error
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

func (i2f *Info2File) Add() error {
	return db.MySQL().Create(i2f).Error
}
