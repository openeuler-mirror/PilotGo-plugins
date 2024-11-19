package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type SSHDFile struct {
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

func (sdf *SSHDFile) Add() error {
	return db.MySQL().Save(&sdf).Error
}

// 根据配置uuid获取用户自己创建的配置信息
func GetSSHDFileByInfoUUID(uuid string, isindex interface{}) (SSHDFile, error) {
	var file SSHDFile
	if isindex != nil {
		err := db.MySQL().Model(&SSHDFile{}).Where("config_info_uuid=? and is_from_host=0 and is_index = ?", uuid, isindex).Find(&file).Error
		return file, err
	}
	err := db.MySQL().Model(&SSHDFile{}).Where("config_info_uuid=? and is_from_host=0", uuid).Find(&file).Error
	return file, err
}

func GetSSHDFileByUUID(uuid string) (SSHDFile, error) {
	var file SSHDFile
	err := db.MySQL().Model(&SSHDFile{}).Where("uuid=?", uuid).Find(&file).Error
	return file, err
}

func (sdf *SSHDFile) UpdateByuuid() error {
	// 将同类配置的所有标志修改为未使用
	err := db.MySQL().Model(&SSHDFile{}).Where("config_info_uuid=?", sdf.ConfigInfoUUID).Update("is_index", 0).Error
	if err != nil {
		return err
	}
	// 将成功下发的具体某一个配置状态修改为已使用
	return db.MySQL().Model(&SSHDFile{}).Where("uuid=?", sdf.UUID).Update("is_index", 1).Error
}

// 根据配置uuid获取所有配置文件
func GetSSHDFilesByConfigUUID(uuid string) ([]SSHDFile, error) {
	var files []SSHDFile
	err := db.MySQL().Model(&SSHDFile{}).Where("config_info_uuid=?", uuid).Find(&files).Error
	return files, err
}
