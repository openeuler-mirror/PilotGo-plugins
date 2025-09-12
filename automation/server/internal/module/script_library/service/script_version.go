package service

import (
	"time"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func GetScriptVersions(scriptId string) (*model.ScriptVersionResponse, error) {
	return dao.GetScriptVersions(scriptId)
}

func AddScriptVersion(sv *model.ScriptVersion) error {
	scriptVersion := &model.ScriptVersion{
		ScriptID:            sv.ScriptID,
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
