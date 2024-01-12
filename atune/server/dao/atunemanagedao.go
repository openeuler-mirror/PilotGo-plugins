package dao

import (
	"errors"

	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func IsExist(uuid string) (bool, error) {
	var ac model.AtuneClient
	err := db.MySQL().Where("machine_uuid = ?", uuid).Find(&ac).Error
	if err != nil {
		return false, errors.New("查询数据库失败：" + err.Error())
	}
	return ac.ID != 0, nil
}

func AddAtuneClientList(ac *model.AtuneClient) error {
	a := model.AtuneClient{
		MachineUUID: ac.MachineUUID,
		MachineIP:   ac.MachineIP,
	}
	err := db.MySQL().Save(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAtuneClientList(ac *model.AtuneClient) error {
	var a model.AtuneClient
	err := db.MySQL().Where("machine_uuid = ?", ac.MachineUUID).Unscoped().Delete(a).Error
	if err != nil {
		return err
	}
	return nil
}
