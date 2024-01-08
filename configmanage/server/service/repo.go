package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/httputils"
	"gitee.com/openeuler/PilotGo/sdk/common"
	"github.com/google/uuid"

	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type RepoConfig struct {
	UUID    string `json:"uuid"` //ConfigInstance的uuid
	Name    string `json:"name"`
	File    string `json:"file"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

func (rc *RepoConfig) Record() error {
	cf := ConfigFile{
		UUID: uuid.New().String(),
		Name: rc.Name,
		File: rc.File,
		Path: rc.Path,
	}
	err := cf.Add()
	if err != nil {
		return err
	}
	i2f := internal.Info2File{
		ConfigInfoUUID: rc.UUID,
		ConfigFileUUID: cf.UUID,
		Version:        rc.Version,
	}
	return i2f.Add()
}

func (c *RepoConfig) Load() error {
	cf, err := internal.GetConfigFileByUUID(c.UUID)
	if err != nil {
		return err
	}
	c.Name = cf.Name
	c.File = cf.File
	c.Path = cf.Path
	return nil
}

func (c *RepoConfig) Apply(de Deploy) (json.RawMessage, error) {
	//TODO:检查de里面的参数是否存在于数据库

	url := "http://" + client.GetClient().Server() + "/api/v1/pluginapi/file_deploy"
	fmt.Println(url)
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
	if resp.Data != nil {
		return resp.Data, errors.New(resp.Message)
	}
	return nil, nil
}

/*
	func (c *RepoConfig) UpdateRepoConfig(configuuid string) error {
		ci, err := internal.GetInfoByConfigUUID(configuuid)
		if err != nil {
			return err
		}
		ci.ConfigFileUUID = c.UUID
		return ci.Add()
	}

	func HistoryRepoConfig(configuuid string) ([]RepoConfig, error) {
		var rcs []RepoConfig
		ci, err := internal.GetInfoByConfigUUID(configuuid)
		if err != nil {
			return nil, err
		}
		cis, err := internal.GetInfoByUUID(ci.UUID)
		for _, v := range cis {
			cf, err := internal.GetConfigFileByUUID(v.ConfigFileUUID)
			if err != nil {
				logger.Error(err.Error())
			}
			rc := RepoConfig{
				UUID: cf.UUID,
				Name: cf.Name,
				File: cf.File,
			}
			rcs = append(rcs, rc)
		}
		return rcs, err
	}
*/
