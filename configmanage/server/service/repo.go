package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo/sdk/utils/httputils"
	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type RepoFile = internal.RepoFile

type RepoConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
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
		Path:           rc.Path,
		Name:           rc.Name,
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
	rc.Path = rf.Path
	rc.Name = rf.Name
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
	Repofiles := common.File{}
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
		Deploy_Text:      Repofiles.Content,
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

func (rc *RepoConfig) Collection() error {
	ci, err := GetConfigByUUID(rc.ConfigInfoUUID)
	if err != nil {
		return err
	}
	//TODO:在nodes中添加批次和部门的机器uuid信息
	nodes := ci.Nodes
	/*for _, v := range ci.BatchIds {
		url := "http://" + client.GetClient().Server() + "/api/v1/pluginapi/batch_uuid?batchId=" + string(v)
		r, err := httputils.Get(url, &httputils.Params{})
		if err != nil {
			return err
		}
		if r.StatusCode != http.StatusOK {
			return errors.New("server process error:" + strconv.Itoa(r.StatusCode))
		}

		resp := &common.CommonResult{}
		if err := json.Unmarshal(r.Body, resp); err != nil {
			return err
		}
		if resp.Code != http.StatusOK {
			return errors.New(resp.Message)
		}
		values := []string{}

	}*/

	//发请求获取配置详情
	url := "http://" + client.GetClient().Server() + "/api/v1/pluginapi/getnodefiles"
	p := struct {
		NodeUUIds []string `json:"nodes"`
		Path      string   `json:"path"`
		FileName  string   `json:"filename"`
	}{
		NodeUUIds: nodes,
		Path:      rc.Path,
		FileName:  rc.Name,
	}
	r, err := httputils.Post(url, &httputils.Params{
		Body: p,
	})
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return errors.New("server process error:" + strconv.Itoa(r.StatusCode))
	}

	resp := &common.CommonResult{}
	if err := json.Unmarshal(r.Body, resp); err != nil {
		return err
	}
	if resp.Code != http.StatusOK {
		return errors.New(resp.Message)
	}
	data := []struct {
		UUID  string
		Error string
		Data  interface{}
	}{}
	if err := resp.ParseData(&data); err != nil {
		return err
	}
	result := ""
	for _, v := range data {
		if v.Error == "" {
			file, _ := json.Marshal(v.Data)
			rf := RepoFile{
				UUID:           uuid.New().String(),
				ConfigInfoUUID: rc.ConfigInfoUUID,
				Content:        file,
				Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
				IsHost:         true,
				Hostuuid:       v.UUID,
			}
			err = rf.Add()
			if err != nil {
				logger.Error(err.Error())
			}
		} else {
			result = result + v.UUID + ":" + v.Error + "\n"
		}
	}
	if result != "" {
		return errors.New(result)
	}
	return nil
}
