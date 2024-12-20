/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Wed Dec 6 15:46:19 2023 +0800
 */
package internal

import "openeuler.org/PilotGo/configmanage-plugin/db"

type ConfigInfo struct {
	ID          int    `gorm:"autoIncrement:true"`
	UUID        string `gorm:"type:varchar(50);primary_key" json:"uuid"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (cm *ConfigInfo) Add() error {
	sql := `
	INSERT INTO config_info (uuid,type,description) 
	VALUES (?, ?, ?) 
	ON DUPLICATE KEY UPDATE
		uuid = VALUES(uuid),
		type = VALUES(type),
		description = VALUES(description);
	`
	return db.MySQL().Exec(sql,
		cm.UUID,
		cm.Type,
		cm.Description,
	).Error
}

func GetInfoByUUID(uuid string) (ConfigInfo, error) {
	var ci ConfigInfo
	err := db.MySQL().Model(&ConfigInfo{}).Where("uuid=?", uuid).Find(&ci).Error
	return ci, err
}

func GetInfos(offset, size int) (int, []*ConfigInfo, error) {
	infos := []*ConfigInfo{}
	var count int64
	err := db.MySQL().Model(&ConfigInfo{}).Count(&count).Error
	if err != nil {
		return 0, infos, err
	}
	err = db.MySQL().Model(&ConfigInfo{}).Limit(size).Offset(offset).Find(&infos).Error
	return int(count), infos, err
}
