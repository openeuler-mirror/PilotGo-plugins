package internal

import (
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type ConfigFile struct {
	ID        int    `gorm:"AUTO_INCREMENT"`
	UUID      string `gorm:"primary_key;type:varchar(50)" json:"uuid"`
	Name      string `json:"name"`
	File      string `gorm:"type:text" json:"file"`
	CreatedAt time.Time
}

func (cf *ConfigFile) Add() error {
	return db.MySQL().Save(&cf).Error
}

func GetConfigFileByUUID(uuid string) (ConfigFile, error) {
	var file ConfigFile
	err := db.MySQL().Where("uuid=?", uuid).Find(&file).Error
	return file, err
}
