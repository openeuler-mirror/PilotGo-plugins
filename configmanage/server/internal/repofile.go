package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type RepoFile struct {
	ID             int        `gorm:"unique;autoIncrement:true"`
	UUID           string     `gorm:"primary_key;type:varchar(50)" json:"uuid"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	Content        json.RawMessage `gorm:"type:json" json:"content"`
	Version        string          `gorm:"type:varchar(50)" json:"version"`
	IsActive       bool            `json:"isactive"`
	IsFromHost     bool            `gorm:"default:false" json:"isfromhost"`
	Hostuuid       string          `gorm:"type:varchar(50)" json:"hostuuid"`
	CreatedAt      time.Time
}

func (rf *RepoFile) Add() error {
	return db.MySQL().Save(&rf).Error
}

func GetRepoFileByInfoUUID(uuid string, isindex interface{}) (RepoFile, error) {
	var file RepoFile
	if isindex != nil {
		err := db.MySQL().Where("config_info_uuid=? && is_index = ?", uuid, isindex).Find(&file).Error
		return file, err
	}
	err := db.MySQL().Where("config_info_uuid=?", uuid).Find(&file).Error
	return file, err
}

func GetRepoFileByUUID(uuid string) (RepoFile, error) {
	var file RepoFile
	err := db.MySQL().Where("uuid=?", uuid).Find(&file).Error
	return file, err
}

func (rf *RepoFile) UpdateByuuid() error {
	err := db.MySQL().Table("repo_file").Where("config_info_uuid=?", rf.ConfigInfoUUID).Update("is_index", 0).Error
	if err != nil {
		return err
	}
	return db.MySQL().Table("repo_file").Where("uuid=?", rf.UUID).Update("is_index", 1).Error
}
