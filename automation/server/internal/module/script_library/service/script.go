package service

import (
	"time"

	"github.com/google/uuid"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

func generateScriptId() string {
	return uuid.NewString()
}

func AddScript(s *model.AddScriptDTO) error {
	scriptId := generateScriptId()

	// decodedContent, err := utils.DecodeScriptContent(s.Content)
	// if err != nil {
	// 	return err
	// }

	script := &model.Script{
		ID:          scriptId,
		Name:        s.Script.Name,
		ScriptType:  s.Script.ScriptType,
		Description: s.Script.Description,
		UsageType:   s.Script.UsageType,
		Tag:         s.Script.Tag,
		ModifyUser:  s.Script.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	}

	scriptVersion := &model.ScriptVersion{
		ScriptID: scriptId,
		// Content:  decodedContent,
		Content:     s.FirstVersion.Content,
		Params:      s.FirstVersion.Params,
		Version:     s.FirstVersion.Version,
		VersionDesc: s.FirstVersion.VersionDesc,
		ModifyUser:  s.Script.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	}

	return dao.AddScript(script, scriptVersion)
}

func GetScripts(query *response.PagedQuery) ([]*model.ScriptResponse, int, error) {
	return dao.GetScripts(query)
}

func UpdateScript(s *model.Script) error {
	script := &model.Script{
		Name:        s.Name,
		Description: s.Description,
		Tag:         s.Tag,
		UsageType:   s.UsageType,
		ModifyUser:  s.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	return dao.UpdateScript(s.ID, script)
}

func DeleteScript(id string) error {
	return dao.DeleteScript(id)
}
