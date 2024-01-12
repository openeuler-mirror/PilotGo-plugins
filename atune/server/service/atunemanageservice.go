package service

import (
	"errors"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/model"
)

var ResultOptMsg = []string{"安装成功", "卸载成功"}

const (
	CommandInstall_Type   = "install"
	CommandUninstall_Type = "uninstall"

	CommandInstall_Cmd   = "yum install -y golang-github-prometheus-node_exporter && (echo '安装成功'; systemctl start node_exporter) || echo '安装失败'"
	CommandUninstall_Cmd = "yum remove -y golang-github-prometheus-node_exporter && echo '卸载成功' || echo '卸载失败'"
)

func AtuneManage(res *common.CmdResult, command_type string) error {
	result := &model.AtuneClient{
		MachineUUID: res.MachineUUID,
		MachineIP:   res.MachineIP,
	}
	logger.Info("A-Tune客户端安装状态:\n%v", res)
	ok, err := dao.IsExist(res.MachineUUID)
	if err != nil {
		return err
	}
	if !ok && command_type == CommandInstall_Type && resultOptStdout(res) {
		if err := dao.AddAtuneClientList(result); err != nil {
			return errors.New("保存结果失败：" + err.Error())
		}
	}
	if ok && command_type == CommandUninstall_Type && resultOptStdout(res) {
		if err := dao.DeleteAtuneClientList(result); err != nil {
			return errors.New("删除prometheus target失败: " + err.Error())
		}
	}

	return nil
}

func resultOptStdout(res *common.CmdResult) bool {
	stdout := res.Stdout
	for _, msg := range ResultOptMsg {
		if strings.Contains(stdout, msg) {
			return true
		}
	}
	return false
}
