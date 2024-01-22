package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type RepoFile struct {
	ID             int        `gorm:"AUTO_INCREMENT"`
	UUID           string     `gorm:"primary_key;type:varchar(50)" json:"uuid"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	Content        json.RawMessage `gorm:"type:json" json:"content"`
	Version        string          `gorm:"type:varchar(50)" json:"version"`
	IsIndex        bool            `json:"isindex"`
	CreatedAt      time.Time
}

func (rf *RepoFile) Add() error {
	return db.MySQL().Save(&rf).Error
}

func GetRepoFileByInfoUUID(uuid string) (RepoFile, error) {
	var file RepoFile
	err := db.MySQL().Where("config_info_uuid=? && is_index = 1", uuid).Find(&file).Error
	return file, err
}
