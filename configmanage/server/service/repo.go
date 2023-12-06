package service

import (
	"openeuler.org/PilotGo/configmanage-plugin/global"
)

type RepoConfig struct {
	ConfigInstance ConfigInstance
	Name           string
	File           string
}

func (c *RepoConfig) Record() error {
	c.ConfigInstance.Type = global.Repo
	uuid, err := c.ConfigInstance.Record()
	if err != nil {
		return err
	}
	cf := ConfigFile{
		ConfigInfoUUID: uuid,
		Name:           c.Name,
		File:           c.File,
	}
	err = cf.Add()
	return err
}

func (c *RepoConfig) Load() error {
	return nil
}

func (c *RepoConfig) Apply(uuid string) error {

	return nil
}
