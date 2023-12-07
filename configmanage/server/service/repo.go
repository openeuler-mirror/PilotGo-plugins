package service

import (
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

type RepoConfig struct {
	UUID string
	Name string
	File string
}

func (c *RepoConfig) Record() error {
	cf := ConfigFile{
		UUID: c.UUID,
		Name: c.Name,
		File: c.File,
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
	return nil
}

func (c *RepoConfig) Apply(uuid string) error {

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
