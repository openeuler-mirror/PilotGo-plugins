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
	IsActive bool `json:"isactive"`
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
		IsActive:       rc.IsActive,
	}
	return rf.Add()
}

func (rc *RepoConfig) Load() error {
	// 加载正在使用的某配置文件
	rf, err := internal.GetRepoFileByInfoUUID(rc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	rc.UUID = rf.UUID
	rc.Path = rf.Path
	rc.Name = rf.Name
	rc.Content = rf.Content
	rc.Version = rf.Version
	rc.IsActive = rf.IsActive
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
	Repofiles := []common.File{}
	err = json.Unmarshal([]byte(rf.Content), &Repofiles)
	if err != nil {
		return nil, err
	}
	result := ""
	for _, v := range Repofiles {
		de := Deploy{
			DeployBatch: common.Batch{
				BatchIds:      batchids,
				DepartmentIDs: departids,
				MachineUUIDs:  nodes,
			},
			DeployPath:     v.Path,
			DeployFileName: v.Name,
			DeployText:     v.Content,
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
				result = result + v.Content + "文件" + d.UUID + ":" + d.Error + "\n"
			}
		}
	}
	//TODO:部分成功如何修改数据库
	if result == "" {
		//下发成功修改数据库应用版本
		err = rf.UpdateByuuid()
		return nil, err
	}
	return nil, errors.New(result + "failed to apply repo config")
}

func (rc *RepoConfig) Collect() error {
	ci, err := GetConfigByUUID(rc.ConfigInfoUUID)
	if err != nil {
		return err
	}

	//发请求获取配置详情
	url := "http://" + client.GetClient().Server() + "/api/v1/pluginapi/getnodefiles"
	p := struct {
		DeployBatch common.Batch `json:"deploybatch"`
		Path        string       `json:"path"`
		FileName    string       `json:"filename"`
	}{
		DeployBatch: common.Batch{
			BatchIds:      ci.BatchIds,
			DepartmentIDs: ci.DepartIds,
			MachineUUIDs:  ci.Nodes,
		},
		Path:     rc.Path,
		FileName: rc.Name,
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

	data := []common.NodeResult{}
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
				IsFromHost:     true,
				Hostuuid:       v.UUID,
			}
			err = rf.Add()
			if err != nil {
				logger.Error("failed to add repoconfig: %s", err.Error())
			}
		} else {
			result = result + v.UUID + ":" + v.Error + "\n"
		}
	}
	if result != "" {
		return errors.New(result + "failed to collect repo config")
	}
	return nil
}

// 根据配置uuid获取所有配置文件
func GetRopeFilesByCinfigUUID(uuid string) ([]RepoFile, error) {
	return internal.GetRopeFilesByCinfigUUID(uuid)
}

// 查看某台机器某种类型的的历史配置信息
func GetRopeFilesByNode(nodeid string) ([]RepoConfig, error) {
	// 查找本台机器所属的配置uuid
	config_nodes, err := internal.GetConfigNodesByNode(nodeid)
	if err != nil {
		return nil, err
	}
	var rcs []RepoConfig
	for _, v := range config_nodes {
		rf, err := internal.GetRepoFileByInfoUUID(v.ConfigInfoUUID, nil)
		if err != nil {
			return nil, err
		}
		rc := RepoConfig{
			UUID:           rf.UUID,
			ConfigInfoUUID: rf.ConfigInfoUUID,
			Path:           rf.Path,
			Name:           rf.Name,
			Content:        rf.Content,
			Version:        rf.Version,
			IsActive:       rf.IsActive,
		}
		rcs = append(rcs, rc)
	}
	return rcs, nil
}
