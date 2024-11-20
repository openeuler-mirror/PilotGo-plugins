package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type SSHFile struct {
	ID             int        `gorm:"autoIncrement:true"`
	UUID           string     `gorm:"primary_key;type:varchar(50)" json:"uuid"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	Content        json.RawMessage `gorm:"type:json" json:"content"`
	Version        string          `gorm:"type:varchar(50)" json:"version"`
	IsActive       bool            `gorm:"default:false" json:"isactive"`
	IsFromHost     bool            `gorm:"default:false" json:"isfromhost"`
	Hostuuid       string          `gorm:"type:varchar(50)" json:"hostuuid"`
	CreatedAt      time.Time
}

func (sf *SSHFile) Add() error {
	return db.MySQL().Save(&sf).Error
}

// 根据配置uuid获取用户自己创建的配置信息
func GetSSHFileByInfoUUID(uuid string, isindex interface{}) (SSHFile, error) {
	var file SSHFile
	if isindex != nil {
		err := db.MySQL().Model(&SSHFile{}).Where("config_info_uuid=? and is_from_host=0 and is_index = ?", uuid, isindex).Find(&file).Error
		return file, err
	}
	err := db.MySQL().Model(&SSHFile{}).Where("config_info_uuid=? and is_from_host=0", uuid).Find(&file).Error
	return file, err
}

func GetSSHFileByUUID(uuid string) (SSHFile, error) {
	var file SSHFile
	err := db.MySQL().Model(&SSHFile{}).Where("uuid=?", uuid).Find(&file).Error
	return file, err
}

func (sf *SSHFile) UpdateByuuid() error {
	// 将同类配置的所有标志修改为未使用
	err := db.MySQL().Model(&SSHFile{}).Where("config_info_uuid=?", sf.ConfigInfoUUID).Update("is_index", 0).Error
	if err != nil {
		return err
	}
	// 将成功下发的具体某一个配置状态修改为已使用
	return db.MySQL().Model(&SSHFile{}).Where("uuid=?", sf.UUID).Update("is_index", 1).Error
}

// 根据配置uuid获取所有配置文件
func GetSSHFilesByConfigUUID(uuid string) ([]SSHFile, error) {
	var files []SSHFile
	err := db.MySQL().Model(&SSHFile{}).Where("config_info_uuid=?", uuid).Find(&files).Error
	return files, err
}
