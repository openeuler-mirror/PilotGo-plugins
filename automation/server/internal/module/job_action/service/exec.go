package service

import (
	"github.com/google/uuid"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/exec"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
)

func ExecScript(ips []string, scriptID, params string, timeoutSec int) error {
	scriptType, encodeScriptContent, err := dao.GetPublishedScriptByScriptId(scriptID)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		sr := exec.ScriptsRun{
			JobId:         uuid.New().String(),
			ScriptType:    scriptType,
			ScriptContent: encodeScriptContent,
			Params:        params,
			TimeOut:       timeoutSec,
		}
		exec.ExecScript(ip, sr)
	}
	return nil
}
