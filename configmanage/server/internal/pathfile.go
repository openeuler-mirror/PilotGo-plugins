/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Tue Nov 26 15:36:42 2024 +0800
 */
package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type PathFile struct {
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

func (pf *PathFile) Add() error {
	sql := `
	INSERT INTO path_file (uuid,config_info_uuid,path,name,content,version,is_active,is_from_host,hostuuid) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) 
	ON DUPLICATE KEY UPDATE
		uuid = VALUES(uuid),
		config_info_uuid = VALUES(config_info_uuid),
		path = VALUES(path),
		name = VALUES(name),
		content = VALUES(content),
		version = VALUES(version),
		is_active = VALUES(is_active),
		is_from_host = VALUES(is_from_host),
		hostuuid = VALUES(hostuuid);
	`
	return db.MySQL().Exec(sql,
		pf.UUID,
		pf.ConfigInfoUUID,
		pf.Path,
		pf.Name,
		pf.Content,
		pf.Version,
		pf.IsActive,
		pf.IsFromHost,
		pf.Hostuuid,
	).Error
}

// 根据配置uuid获取用户自己创建的配置信息
func GetPathFileByInfoUUID(uuid string, isindex interface{}) (PathFile, error) {
	var file PathFile
	if isindex != nil {
		err := db.MySQL().Model(&PathFile{}).Where("config_info_uuid=? and is_from_host=0 and is_index = ?", uuid, isindex).Find(&file).Error
		return file, err
	}
	err := db.MySQL().Model(&PathFile{}).Where("config_info_uuid=? and is_from_host=0 ", uuid).Find(&file).Error
	return file, err
}

func GetPathFileByUUID(uuid string) (PathFile, error) {
	var file PathFile
	err := db.MySQL().Model(&PathFile{}).Where("uuid=?", uuid).Find(&file).Error
	return file, err
}

func (pf *PathFile) UpdateByuuid() error {
	// 将同类配置的所有标志修改为未使用
	err := db.MySQL().Model(&PathFile{}).Where("config_info_uuid=?", pf.ConfigInfoUUID).Update("is_index", 0).Error
	if err != nil {
		return err
	}
	// 将成功下发的具体某一个配置状态修改为已使用
	return db.MySQL().Model(&PathFile{}).Where("uuid=?", pf.UUID).Update("is_index", 1).Error
}

// 根据配置uuid获取所有配置文件
func GetPathFilesByConfigUUID(uuid string) ([]PathFile, error) {
	var files []PathFile
	err := db.MySQL().Model(&PathFile{}).Where("config_info_uuid=?", uuid).Find(&files).Error
	return files, err
}
