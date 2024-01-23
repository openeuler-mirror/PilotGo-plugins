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

type RepoFile = internal.RepoFile

type RepoConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	//下发改变标志位
	IsIndex bool `json:"isindex"`
}

func (rc *RepoConfig) Record() error {
	//检查info的uuid是否存在
	ci, err := GetInfoByUUID(rc.ConfigInfoUUID)
	if err != nil || ci.UUID == "" {
		return errors.New("configinfo uuid not exist")
	}

	rf := RepoFile{
		UUID:           rc.UUID,
		ConfigInfoUUID: rc.ConfigInfoUUID,
		Content:        rc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsIndex:        rc.IsIndex,
	}
	return rf.Add()
}

func (rc *RepoConfig) Load() error {
	rf, err := internal.GetRepoFileByInfoUUID(rc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	rc.UUID = rf.UUID
	rc.Content = rf.Content
	rc.Version = rf.Version
	rc.IsIndex = rf.IsIndex
	return nil
}

func (rc *RepoConfig) Apply() (json.RawMessage, error) {
	//从数据库获取下发的信息
	rf, err := internal.GetRepoFileByUUID(rc.UUID)
	if err != nil {
		return nil, err
	}
	if rf.ConfigInfoUUID != rc.ConfigInfoUUID || rf.UUID != rc.UUID {
		return nil, errors.New("数据库不存在此配置")
	}

	batchids, err := internal.GetConfigBatchByUUID(rc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(rc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	nodes, err := internal.GetConfigNodesByUUID(rc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	//从rc中解析下发的文件内容，逐一进行下发
	Repofiles := struct {
		Name string `json:"name"`
		File string `json:"file"`
		Path string `json:"path"`
	}{}
	err = json.Unmarshal(rf.Content, &Repofiles)
	if err != nil {
		return nil, err
	}
	de := Deploy{
		Deploy_BatchIds:  batchids,
		Deploy_DepartIds: departids,
		Deploy_NodeUUIds: nodes,
		Deploy_Path:      Repofiles.Path,
		Deploy_FileName:  Repofiles.Name,
		Deploy_Text:      Repofiles.File,
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

	if string(resp.Data) != "null" {
		return resp.Data, errors.New(resp.Message)
	}
	//下发成功修改数据库应用版本
	err = rf.UpdateByuuid()
	return nil, err
}
