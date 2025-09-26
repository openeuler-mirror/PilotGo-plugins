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
		ScriptID:    scriptId,
		Content:     sv.Content,
		Params:      sv.Params,
		Version:     sv.Version,
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
		Version:     sv.Version,
		VersionDesc: sv.VersionDesc,
		ModifyUser:  sv.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	return dao.UpdateScriptVersion(sv.ID, scriptId, scriptVersion)
}

func DeleteScriptVersion(id int, scriptId string) error {
	return dao.DeleteScriptVersion(id, scriptId)
}
