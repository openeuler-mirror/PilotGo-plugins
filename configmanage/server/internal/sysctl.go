package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type SysctlFile struct {
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

func (sysf *SysctlFile) Add() error {
	return db.MySQL().Save(&sysf).Error
}

func GetSysctlFileByInfoUUID(uuid string, isindex interface{}) (SysctlFile, error) {
	var file SysctlFile
	if isindex != nil {
		err := db.MySQL().Model(&SysctlFile{}).Where("config_info_uuid=? && is_index = ?", uuid, isindex).Find(&file).Error
		return file, err
	}
	err := db.MySQL().Model(&SysctlFile{}).Where("config_info_uuid=?", uuid).Find(&file).Error
	return file, err
}
