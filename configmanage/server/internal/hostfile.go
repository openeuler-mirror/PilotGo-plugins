package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type HostFile struct {
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

func (hf *HostFile) Add() error {
	return db.MySQL().Save(&hf).Error
}

func GetHostFileByInfoUUID(uuid string, isindex interface{}) (HostFile, error) {
	var file HostFile
	if isindex != nil {
		err := db.MySQL().Model(&HostFile{}).Where("config_info_uuid=? && is_index = ?", uuid, isindex).Find(&file).Error
		return file, err
	}
	err := db.MySQL().Model(&HostFile{}).Where("config_info_uuid=?", uuid).Find(&file).Error
	return file, err
}

func GetHostFileByUUID(uuid string) (HostFile, error) {
	var file HostFile
	err := db.MySQL().Model(&HostFile{}).Where("uuid=?", uuid).Find(&file).Error
	return file, err
}

func (hf *HostFile) UpdateByuuid() error {
	// 将同类配置的所有标志修改为未使用
	err := db.MySQL().Model(&HostFile{}).Where("config_info_uuid=?", hf.ConfigInfoUUID).Update("is_index", 0).Error
	if err != nil {
		return err
	}
	// 将成功下发的具体某一个配置状态修改为已使用
	return db.MySQL().Model(&HostFile{}).Where("uuid=?", hf.UUID).Update("is_index", 1).Error
}

// 根据配置uuid获取所有配置文件
func GetHostFilesByCinfigUUID(uuid string) ([]HostFile, error) {
	var files []HostFile
	err := db.MySQL().Model(&HostFile{}).Where("config_info_uuid=?", uuid).Find(&files).Error
	return files, err
}
