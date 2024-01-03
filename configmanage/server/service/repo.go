package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/httputils"
	"gitee.com/openeuler/PilotGo/sdk/common"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type RepoConfig struct {
	UUID string
	Name string
	File string
	Path string
}

func (c *RepoConfig) Record() error {
	cf := ConfigFile{
		UUID: c.UUID,
		Name: c.Name,
		File: c.File,
		Path: c.Path,
	}
	return cf.Add()
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

func (c *RepoConfig) Apply(de Deploy) ([]string, error) {
	//检查de里面的参数是否存在于数据库

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
	fmt.Println(resp.Data)
	return nil, nil
}
