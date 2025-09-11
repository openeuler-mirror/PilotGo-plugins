package service

import (
	"time"

	"github.com/google/uuid"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func generateScriptId() string {
	return uuid.NewString()
}

func AddScript(s *model.ScriptWithVersion) error {
	scriptId := generateScriptId()

	script := &model.Script{
		ID:                  scriptId,
		Name:                s.Name,
		ScriptType:          s.ScriptType,
		Description:         s.Description,
		Tags:                s.Tags,
		IsPublic:            s.IsPublic,
		Creator:             s.Creator,
		CreatedAt:           time.Now().Format("2006-01-02 15:04:05"),
		LastModifyUser:      s.Creator,
		LastModifyUpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	scriptVersion := &model.ScriptVersion{
		ScriptID:            scriptId,
		Content:             s.Content,
		Version:             s.Version,
		VersionDesc:         s.VersionDesc,
		Creator:             s.Creator,
		CreatedAt:           time.Now().Format("2006-01-02 15:04:05"),
		LastModifyUser:      s.Creator,
		LastModifyUpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return dao.AddScript(script, scriptVersion)
}
