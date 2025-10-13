package service

import (
	"ant-agent/exec-script/model"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/logger"
)

func getScriptInfo(scriptType string) (suffix string, interpreter string, err error) {
	switch scriptType {
	case "Shell":
		return ".sh", "/bin/bash", nil
	case "Python":
		return ".py", "/usr/bin/python3", nil
	case "Perl":
		return ".pl", "/usr/bin/perl", nil
	case "Ruby":
		return ".rb", "/usr/bin/ruby", nil
	case "PHP":
		return ".php", "/usr/bin/php", nil
	default:
		err = fmt.Errorf("不支持的脚本类型: %s", scriptType)
	}
	return
}

func createTempScriptFile(workDir, suffix, encodedScript string) (string, error) {
	decodedScript, err := base64.StdEncoding.DecodeString(encodedScript)
	if err != nil {
		return "", fmt.Errorf("脚本内容base64解码失败: %s", err.Error())
	}

	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		if err := os.MkdirAll(workDir, 0755); err != nil {
			return "", fmt.Errorf("创建临时工作目录失败: %s", err.Error())
		}
	}

	tmpFile, err := os.CreateTemp(workDir, "script_*"+suffix)
	if err != nil {
		return "", fmt.Errorf("创建临时脚本文件失败: %s", err.Error())
	}

	if _, err = tmpFile.Write(decodedScript); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("内容写入到脚本文件失败: %s", err.Error())
	}
	tmpFile.Close()

	if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("设置脚本可执行权限失败: %s", err.Error())
	}

	return tmpFile.Name(), nil
}

func runScript(interpreter, scriptPath string, params string, timeoutSec int) (string, string, int, error) {
	if timeoutSec <= 0 {
		timeoutSec = 30
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	args := append([]string{scriptPath}, params)
	cmd := exec.CommandContext(ctx, interpreter, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return stdout.String(), "脚本执行超时", -1, fmt.Errorf("脚本执行超时")
	}

	if err != nil {
		retcode := -1
		if cmd.ProcessState != nil {
			retcode = cmd.ProcessState.ExitCode()
		}
		errMsg := stderr.String()
		if errMsg == "" {
			errMsg = err.Error()
		}
		return stdout.String(), errMsg, retcode, err
	}

	return stdout.String(), stderr.String(), 0, nil
}

func ExecScript(script *model.ScriptsRun) (interface{}, error) {
	var workDir = "/tmp/scripts/"

	suffix, interpreter, err := getScriptInfo(script.ScriptType)
	if err != nil {
		return "", fmt.Errorf("获取脚本类型失败: %s", err.Error())
	}

	scriptPath, err := createTempScriptFile(workDir, suffix, script.ScriptContent)
	if err != nil {
		return "", fmt.Errorf("创建临时脚本失败: %s", err.Error())
	}
	defer os.Remove(scriptPath)

	logger.Debug("run script timeout: %v", script.TimeOut)
	logger.Debug("process run script command: %s %s %v", interpreter, scriptPath, script.Params)

	stdout, stderr, execCode, err := runScript(interpreter, scriptPath, script.Params, script.TimeOut)
	result := &model.CmdResult{
		Stdout:  stdout,
		Stderr:  stderr,
		RetCode: execCode,
	}
	if err != nil || execCode != 0 {
		return "", fmt.Errorf("脚本执行错误: %s", err.Error())
	}

	return result, nil
}
