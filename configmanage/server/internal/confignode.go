/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Mon Dec 4 16:47:30 2023 +0800
 */
package internal

import (
	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type ConfigNode struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	NodeId         string `json:"node_id"` //机器uuid
}

func (cn *ConfigNode) Add() error {
	return db.MySQL().Create(&cn).Error
}

func GetConfigNodesByUUID(uuid string) ([]string, error) {
	var nodes []string
	err := db.MySQL().Model(&ConfigNode{}).Where("config_info_uuid=?", uuid).Select("node_id").Find(&nodes).Error
	return nodes, err
}

func GetConfigNodesByNode(uuid string) ([]ConfigNode, error) {
	var cns []ConfigNode
	err := db.MySQL().Model(&ConfigNode{}).Where("node_id=?", uuid).Find(&cns).Error
	return cns, err
}

func DelConfigNodeByUUID(uuid string) error {
	err := db.MySQL().Model(&ConfigNode{}).Where("config_info_uuid=?", uuid).Delete(&ConfigNode{}).Error
	return err
}

type ConfigBatch struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	BatchID        int `json:"batch_id"`
}

func (cb *ConfigBatch) Add() error {
	return db.MySQL().Create(&cb).Error
}

func GetConfigBatchByUUID(uuid string) ([]int, error) {
	var nodes []int
	err := db.MySQL().Model(&ConfigBatch{}).Where("config_info_uuid=?", uuid).Select("batch_id").Find(&nodes).Error
	return nodes, err
}

func DelConfigBatchByUUID(uuid string) error {
	err := db.MySQL().Where("config_info_uuid=?", uuid).Delete(&ConfigBatch{}).Error
	return err
}

type ConfigDepart struct {
	ID             int        `gorm:"primary_key;AUTO_INCREMENT"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	DepartID       int `json:"depart_id"`
}

func (cd *ConfigDepart) Add() error {
	return db.MySQL().Create(&cd).Error
}

func GetConfigDepartByUUID(uuid string) ([]int, error) {
	var nodes []int
	err := db.MySQL().Model(&ConfigDepart{}).Where("config_info_uuid=?", uuid).Select("depart_id").Find(&nodes).Error
	return nodes, err
}

func DelConfigDepartByUUID(uuid string) error {
	err := db.MySQL().Where("config_info_uuid=?", uuid).Delete(&ConfigDepart{}).Error
	return err
}
