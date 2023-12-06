package internal

import (
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type ConfigFile struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	Name           string `json:"name"`
	File           string `gorm:"type:text" json:"file"`
	CreatedAt      time.Time
}

func (cf *ConfigFile) Add() error {
	return db.MySQL().Save(&cf).Error
}

func GetConfigFilesByUUID(uuid string) ([]ConfigFile, error) {
	var files []ConfigFile
	err := db.MySQL().Where("config_message_uuid=?", uuid).Find(&files).Error
	return files, err
}
