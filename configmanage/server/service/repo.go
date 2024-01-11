package service

import (
	"fmt"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type RepoConfig struct {
	UUID           string      `json:"uuid"`
	ConfigInfoUUID string      `json:"configinfouuid"` //ConfigInstance的uuid
	Content        interface{} `json:"content"`
	Version        string      `json:"version"`
	IsIndex        bool        `json:"isindex"`
}

func (rc *RepoConfig) Record() error {
	//TODO:检查info的uuid是否存在
	rf := RepoFile{
		UUID:           rc.UUID,
		ConfigInfoUUID: rc.ConfigInfoUUID,
		Content:        rc.Content,
		Version:        fmt.Sprintf("v%s", time.Now().Format("2006-01-02-15-04-05")),
		IsIndex:        rc.IsIndex,
	}
	return rf.Add()
}

func (c *RepoConfig) Load() error {
	rf, err := internal.GetRepoFileByUUID(c.UUID)
	if err != nil {
		return err
	}
	c.UUID = rf.UUID
	c.Content = rf.Content
	c.Version = rf.Version
	c.IsIndex = rf.IsIndex
	return nil
}

/*
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
*/
