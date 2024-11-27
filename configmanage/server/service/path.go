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
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

/*
path: 环境变量文件只有一个

一般方法：	1、在/etc/profile中修改内容

考虑的问题：
*/

type PathFile = internal.PathFile

type PathConfig struct {
	UUID           string          `json:"uuid"`
	ConfigInfoUUID string          `json:"configinfouuid"`
	Content        json.RawMessage `json:"content"`
	Version        string          `json:"version"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	//下发改变标志位
	IsActive bool `json:"isactive"`
}

func (pc *PathConfig) toPathFile() PathFile {
	return PathFile{
		UUID:           pc.UUID,
		ConfigInfoUUID: pc.ConfigInfoUUID,
		Path:           pc.Path,
		Name:           pc.Name,
		Content:        pc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsActive:       pc.IsActive,
		IsFromHost:     false,
	}
}

func toPathConfig(pf *PathFile) PathConfig {
	return PathConfig{
		UUID:           pf.UUID,
		ConfigInfoUUID: pf.ConfigInfoUUID,
		Path:           pf.Path,
		Name:           pf.Name,
		Content:        pf.Content,
		Version:        pf.Version,
		IsActive:       pf.IsActive,
	}
}

func (pc *PathConfig) Record() error {
	pf := pc.toPathFile()
	return pf.Add()
}

func (pc *PathConfig) Load() error {
	// 加载正在使用的某配置文件
	pf, err := internal.GetPathFileByInfoUUID(pc.ConfigInfoUUID, true)
	if err != nil {
		return err
	}
	pc.UUID = pf.UUID
	pc.Path = pf.Path
	pc.Name = pf.Name
	pc.Content = pf.Content
	pc.Version = pf.Version
	pc.IsActive = pf.IsActive
	return nil
}

func (pc *PathConfig) Apply() ([]NodeResult, error) {
	//从数据库获取下发的信息
	pf, err := internal.GetPathFileByUUID(pc.UUID)
	if err != nil {
		return nil, err
	}
	if pf.ConfigInfoUUID != pc.ConfigInfoUUID || pf.UUID != pc.UUID {
		return nil, errors.New("数据库不存在此配置")
	}

	batchids, err := internal.GetConfigBatchByUUID(pc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	departids, err := internal.GetConfigDepartByUUID(pc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}
	nodes, err := internal.GetConfigNodesByUUID(pc.ConfigInfoUUID)
	if err != nil {
		return nil, err
	}

	//从pc中解析下发的文件内容，逐一进行下发
	pathfile := common.File{}
	err = json.Unmarshal([]byte(pf.Content), &pathfile)
	if err != nil {
		return nil, err
	}
	results := []NodeResult{}
	de := Deploy{
		DeployBatch: common.Batch{
			BatchIds:      batchids,
			DepartmentIDs: departids,
			MachineUUIDs:  nodes,
		},
		DeployPath:     pathfile.Path,
		DeployFileName: pathfile.Name,
		DeployText:     pathfile.Content,
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
		// 存储每一台机器的执行结果
		pfNode := PathFile{
			UUID:           pf.UUID,
			ConfigInfoUUID: pf.ConfigInfoUUID,
			Path:           pf.Path,
			Name:           pf.Name,
			Content:        pf.Content,
			Version:        pf.Version,
			IsActive:       true,
			IsFromHost:     false,
			Hostuuid:       d.UUID,
			CreatedAt:      time.Now(),
		}

		// 返回执行失败的机器详情
		if d.Error != "" {
			pfNode.IsActive = false
			results = append(results, NodeResult{
				Type:     global.PATH,
				NodeUUID: d.UUID,
				Detail:   pathfile.Content,
				Result:   false,
				Err:      d.Error,
			})
		}
		err = pfNode.Add()
		if err != nil {
			results = append(results, NodeResult{
				Type:     global.PATH,
				NodeUUID: d.UUID,
				Detail:   "failed to collect path config to db",
				Result:   false,
				Err:      err.Error(),
			})
		}
	}

	// 全部下发成功直接修改数据库是否激活字段
	if results == nil {
		//下发成功修改数据库应用版本
		err = pf.UpdateByuuid()
		return nil, err
	}
	return results, errors.New("failed to apply path config")
}

func (pc *PathConfig) Collect() ([]NodeResult, error) {
	ci, err := GetConfigByUUID(pc.ConfigInfoUUID)
	if err != nil {
		return nil, err
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
		Path:     pc.Path,
		FileName: pc.Name,
	}
	r, err := httputils.Post(url, &httputils.Params{
		Body: p,
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
	results := []NodeResult{}
	for _, v := range data {
		if v.Error == "" {
			file, _ := json.Marshal(v.Data)
			pf := PathFile{
				UUID:           uuid.New().String(),
				ConfigInfoUUID: pc.ConfigInfoUUID,
				Path:           pc.Path,
				Name:           pc.Name,
				Content:        file,
				Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
				IsActive:       true,
				IsFromHost:     true,
				Hostuuid:       v.UUID,
				CreatedAt:      time.Now(),
			}
			err = pf.Add()
			if err != nil {
				logger.Error("failed to add pathconfig: %s", err.Error())
				results = append(results, NodeResult{
					Type:     global.PATH,
					NodeUUID: v.UUID,
					Detail:   "failed to collect path config to db",
					Result:   false,
					Err:      err.Error()})
			}
		} else {
			results = append(results, NodeResult{
				Type:     global.PATH,
				NodeUUID: v.UUID,
				Detail:   "failed to collect path config:" + v.Data.(string),
				Result:   false,
				Err:      v.Error})
		}
	}
	if results != nil {
		return results, errors.New("failed to collect path config")
	}
	return nil, nil
}

func GetPathFileByInfoUUID(uuid string, isindex interface{}) (PathFile, error) {
	return internal.GetPathFileByInfoUUID(uuid, isindex)
}

// 根据配置uuid获取所有配置文件
func GetPathFilesByConfigUUID(uuid string) ([]PathFile, error) {
	return internal.GetPathFilesByConfigUUID(uuid)
}

// 查看某台机器某种类型的的历史配置信息
func GetPathFilesByNode(nodeid string) ([]PathConfig, error) {
	// 查找本台机器所属的配置uuid
	config_nodes, err := internal.GetConfigNodesByNode(nodeid)
	if err != nil {
		return nil, err
	}
	var pcs []PathConfig
	for _, v := range config_nodes {
		pf, err := internal.GetPathFileByInfoUUID(v.ConfigInfoUUID, nil)
		if err != nil {
			return nil, err
		}
		pc := toPathConfig(&pf)
		pcs = append(pcs, pc)
	}
	return pcs, nil
}
