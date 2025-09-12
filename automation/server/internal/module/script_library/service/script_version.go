package service

import (
	"time"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func GetScriptVersions(scriptId string) (*model.ScriptVersionResponse, error) {
	return dao.GetScriptVersions(scriptId)
}

func AddScriptVersion(scriptId string, sv *model.ScriptVersion) error {
	scriptVersion := &model.ScriptVersion{
		ScriptID:            scriptId,
		Content:             sv.Content,
		Version:             sv.Version,
		VersionDesc:         sv.VersionDesc,
		Creator:             sv.Creator,
		CreatedAt:           time.Now().Format("2006-01-02 15:04:05"),
		LastModifyUser:      sv.Creator,
		LastModifyUpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	return dao.AddScriptVersion(scriptVersion)
}

func UpdateScriptVersion(scriptId string, sv *model.ScriptVersion) error {
	scriptVersion := &model.ScriptVersion{
		Content:             sv.Content,
		Version:             sv.Version,
		VersionDesc:         sv.VersionDesc,
		LastModifyUser:      sv.LastModifyUser,
		LastModifyUpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	return dao.UpdateScriptVersion(sv.ID, scriptId, scriptVersion)
}

func DeleteScriptVersion(id int, scriptId string) error {
	return dao.DeleteScriptVersion(id, scriptId)
}
