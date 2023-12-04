package internal

import (
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type ConfigFile struct {
	ID                int           `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigMessage     ConfigMessage `gorm:"Foreignkey:ConfigMessageUUID"`
	ConfigMessageUUID string
	Name              string `json:"name"`
	File              string `gorm:"type:text" json:"file"`
	CreatedAt         time.Time
}

func (cf *ConfigFile) AddConfigFile() error {
	return db.MySQL().Save(&cf).Error
}

func GetConfigFilesByUUID(uuid string) ([]ConfigFile, error) {
	var files []ConfigFile
	err := db.MySQL().Where("config_message_uuid=?", uuid).Find(&files).Error
	return files, err
}
