package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
sshd: 配置文件，sshd_config文件只有一个

一般方法：	1、在/etc/ssh/sshd_config中修改内容

	2、执行命令systemctl restart sshd重启sshd服务

考虑的问题：
*/
type SSHDFile = internal.SSHDFile
type SSHDConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (sdc *SSHDConfig) toSSHDFile() SSHDFile {
	return SSHDFile{
		UUID:           sdc.UUID,
		ConfigInfoUUID: sdc.ConfigInfoUUID,
		Path:           sdc.Path,
		Name:           sdc.Name,
		Content:        sdc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       sdc.IsActive,
	}
}

func toSSHDConfig(sdf *SSHDFile) SSHDConfig {
	return SSHDConfig{
		UUID:           sdf.UUID,
		ConfigInfoUUID: sdf.ConfigInfoUUID,
		Path:           sdf.Path,
		Name:           sdf.Name,
		Content:        sdf.Content,
		Version:        sdf.Version,
		IsActive:       sdf.IsActive,
	}
}

func (sdc *SSHDConfig) Record() error {
	//检查info的uuid是否存在
	ci, err := GetInfoByUUID(sdc.ConfigInfoUUID)
	if err != nil || ci.UUID == "" {
		return errors.New("configinfo uuid not exist")
	}

	sdf := sdc.toSSHDFile()
	return sdf.Add()
}

func (sdc *SSHDConfig) Load() error {
	// 加载正在使用的某配置文件
	sdf, err := internal.GetSSHFileByInfoUUID(sdc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	sdc.UUID = sdf.UUID
	sdc.Path = sdf.Path
	sdc.Name = sdf.Name
	sdc.Content = sdf.Content
	sdc.Version = sdf.Version
	sdc.IsActive = sdf.IsActive
	return nil
}

// TODO:
func (sdc *SSHDConfig) Apply() (json.RawMessage, error) {
	// 从数据库获取下发的信息
	sdf, err := internal.GetSSHDFileByUUID(sdc.UUID)
	if err != nil {
		return nil, err
	}
	if sdf.ConfigInfoUUID != sdc.ConfigInfoUUID || sdf.UUID != sdc.UUID {
		return nil, errors.New("数据库不存在此配置")
	}

	batchids, err := internal.GetConfigBatchByUUID(sdc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(sdc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	nodes, err := internal.GetConfigNodesByUUID(sdc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	// 从hc中解析下发的文件内容，逐一进行下发
	Repofile := common.File{}
	err = json.Unmarshal([]byte(sdf.Content), &Repofile)
	if err != nil {
		return nil, err
	}
	result := ""
	de := Deploy{
		DeployBatch: common.Batch{
			BatchIds:      batchids,
			DepartmentIDs: departids,
			MachineUUIDs:  nodes,
		},
		DeployPath:     Repofile.Path,
		DeployFileName: Repofile.Name,
		DeployText:     Repofile.Content,
	}
	url := "http://" + client.GetClient().Server() + "/api/v1/pluginapi/file_deploy"
	r, err := httputils.Post(url, &httputils.Params{
		Body: de,
	})
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, errors.New("server process error:" + strconv.Itoa(r.StatusCode))
	}

	resp := &common.CommonResult{}
	if err := json.Unmarshal(r.Body, resp); err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.New(resp.Message)
	}

	data := []common.NodeResult{}
	if err := resp.ParseData(&data); err != nil {
		return nil, err
	}
	// 将执行失败的文件、机器信息和原因添加到结果字符串中
	for _, d := range data {
		if d.Error != "" {
			result = result + Repofile.Content + "文件" + d.UUID + ":" + d.Error + "\n"
		}
	}

	// TODO:部分成功如何修改数据库
	if result == "" {
		//下发成功修改数据库应用版本
		err = sdf.UpdateByuuid()
		return nil, err
	}
	return nil, errors.New(result + "failed to apply SSHDConfig")
}

// TODO:
func (sdc *SSHDConfig) Collect() error {
	return nil
}

// 根据配置uuid获取所有配置文件
func GetSSHDFilesByCinfigUUID(uuid string) ([]SSHDFile, error) {
	return internal.GetSSHDFilesByCinfigUUID(uuid)
}

// 查看某台机器某种类型的的历史配置信息
func GetSSHDFilesByNode(nodeid string) ([]SSHDConfig, error) {
	// 查找本台机器所属的配置uuid
	config_nodes, err := internal.GetConfigNodesByNode(nodeid)
	if err != nil {
		return nil, err
	}
	var sdcs []SSHDConfig
	for _, v := range config_nodes {
		sdf, err := internal.GetSSHDFileByInfoUUID(v.ConfigInfoUUID, nil)
		if err != nil {
			return nil, err
		}
		sdc := toSSHDConfig(&sdf)
		sdcs = append(sdcs, sdc)
	}
	return sdcs, nil
}
