package service

import (
	"fmt"
	"strconv"
	"time"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func GetScriptVersions(scriptId string) (*model.ScriptVersionResponse, error) {
	return dao.GetScriptVersions(scriptId)
}

func AddScriptVersion(scriptId string, sv *model.ScriptVersion) error {
	scriptVersion := &model.ScriptVersion{
		ScriptID:    scriptId,
		Content:     sv.Content,
		Params:      sv.Params,
		Version:     nextVersion(scriptId),
		VersionDesc: sv.VersionDesc,
		ModifyUser:  sv.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	return dao.AddScriptVersion(scriptVersion)
}

func UpdateScriptVersion(scriptId string, sv *model.ScriptVersion) error {
	scriptVersion := &model.ScriptVersion{
		Content:     sv.Content,
		Params:      sv.Params,
		VersionDesc: sv.VersionDesc,
		ModifyUser:  sv.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	return dao.UpdateScriptVersion(sv.ID, scriptId, scriptVersion)
}

func DeleteScriptVersion(id int, scriptId string) error {
	return dao.DeleteScriptVersion(id, scriptId)
}

func PublishScriptVersion(id int, scriptId string, newStatus string) error {
	return dao.PublishScriptVersion(id, scriptId, newStatus)
}

func nextVersion(scriptId string) string {
	currentVersion, err := dao.GetLatestScriptVersion(scriptId)
	if err != nil {
		return generateFirstVersion()
	}
	versionNum, _ := strconv.Atoi(currentVersion[1:])
	return fmt.Sprintf("V%d", versionNum+1)
}
