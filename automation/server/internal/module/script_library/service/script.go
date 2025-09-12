package service

import (
	"time"

	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func generateScriptId() string {
	return uuid.NewString()
}

func AddScript(s *model.ScriptWithVersion) error {
	scriptId := generateScriptId()

	// decodedContent, err := utils.DecodeScriptContent(s.Content)
	// if err != nil {
	// 	return err
	// }

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
		ScriptID: scriptId,
		// Content:  decodedContent,
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

func GetScripts(query *response.PaginationQ) ([]*model.ScriptResponse, int, error) {
	return dao.GetScripts(query)
}

func UpdateScript(s *model.Script) error {
	script := &model.Script{
		Name:                s.Name,
		Description:         s.Description,
		Tags:                s.Tags,
		LastModifyUser:      s.LastModifyUser,
		LastModifyUpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	return dao.UpdateScript(s.ID, script)
}

func DeleteScript(id string) error {
	return dao.DeleteScript(id)
}
