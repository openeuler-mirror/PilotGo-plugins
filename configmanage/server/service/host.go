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
host: 配置文件，hosts文件只有一个

一般方法：	1、在/etc/hosts中修改内容

	2、执行命令systemctl restart NetworkManager重启网络服务

考虑的问题：
*/
type HostFile = internal.HostFile

type HostConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (hc *HostConfig) toHostFile() HostFile {
	return HostFile{
		UUID:           hc.UUID,
		ConfigInfoUUID: hc.ConfigInfoUUID,
		Path:           hc.Path,
		Name:           hc.Name,
		Content:        hc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       hc.IsActive,
	}
}

func (hc *HostConfig) Record() error {
	//检查info的uuid是否存在
	ci, err := GetInfoByUUID(hc.ConfigInfoUUID)
	if err != nil || ci.UUID == "" {
		return errors.New("configinfo uuid not exist")
	}

	hf := hc.toHostFile()
	return hf.Add()
}

func (hc *HostConfig) Load() error {
	// 加载正在使用的某配置文件
	hf, err := internal.GetHostFileByInfoUUID(hc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	hc.UUID = hf.UUID
	hc.Path = hf.Path
	hc.Name = hf.Name
	hc.Content = hf.Content
	hc.Version = hf.Version
	hc.IsActive = hf.IsActive
	return nil
}

func (hc *HostConfig) Apply() (json.RawMessage, error) {
	//从数据库获取下发的信息
	hf, err := internal.GetHostFileByUUID(hc.UUID)
	if err != nil {
		return nil, err
	}
	if hf.ConfigInfoUUID != hc.ConfigInfoUUID || hf.UUID != hc.UUID {
		return nil, errors.New("数据库不存在此配置")
	}

	batchids, err := internal.GetConfigBatchByUUID(hc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(hc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	nodes, err := internal.GetConfigNodesByUUID(hc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	//从hc中解析下发的文件内容，逐一进行下发
	Repofile := common.File{}
	err = json.Unmarshal([]byte(hf.Content), &Repofile)
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

	//TODO:部分成功如何修改数据库
	if result == "" {
		//下发成功修改数据库应用版本
		err = hf.UpdateByuuid()
		return nil, err
	}
	return nil, errors.New(result + "failed to apply host config")
}

// TODO:
func (hc *HostConfig) Collect() error {
	return nil
}

// 根据配置uuid获取所有配置文件
func GetHostFilesByConfigUUID(uuid string) ([]HostFile, error) {
	return internal.GetHostFilesByConfigUUID(uuid)
}