/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Wed Nov 15 16:30:21 2023 +0800
 */
package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type SysctlFile struct {
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

func (sysf *SysctlFile) Add() error {
	return db.MySQL().Save(&sysf).Error
}

// 根据配置uuid获取用户自己创建的配置信息
func GetSysctlFileByInfoUUID(uuid string, isindex interface{}) (SysctlFile, error) {
	var file SysctlFile
	if isindex != nil {
		err := db.MySQL().Model(&SysctlFile{}).Where("config_info_uuid=? and is_from_host=0 and is_index = ?", uuid, isindex).Find(&file).Error
		return file, err
	}
	err := db.MySQL().Model(&SysctlFile{}).Where("config_info_uuid=? and is_from_host=0", uuid).Find(&file).Error
	return file, err
}

func GetSysctlFileByUUID(uuid string) (SysctlFile, error) {
	var file SysctlFile
	err := db.MySQL().Model(&SysctlFile{}).Where("uuid=?", uuid).Find(&file).Error
	return file, err
}

func (sysf *SysctlFile) UpdateByuuid() error {
	// 将同类配置的所有标志修改为未使用
	err := db.MySQL().Model(&SysctlFile{}).Where("config_info_uuid=?", sysf.ConfigInfoUUID).Update("is_index", 0).Error
	if err != nil {
		return err
	}
	// 将成功下发的具体某一个配置状态修改为已使用
	return db.MySQL().Model(&SysctlFile{}).Where("uuid=?", sysf.UUID).Update("is_index", 1).Error
}

// 根据配置uuid获取所有配置文件
func GetSysctlFilesByConfigUUID(uuid string) ([]SysctlFile, error) {
	var files []SysctlFile
	err := db.MySQL().Model(&SysctlFile{}).Where("config_info_uuid=?", uuid).Find(&files).Error
	return files, err
}
