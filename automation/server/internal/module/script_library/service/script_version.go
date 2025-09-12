package service

import (
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func GetScriptVersions(scriptId string) (*model.ScriptVersionResponse, error) {
	return dao.GetScriptVersions(scriptId)
}
