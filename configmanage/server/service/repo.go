package service

type RepoConfig struct {
	UUID string
	Name string
	File string
}

func (c *RepoConfig) Record() error {
	cf := ConfigFile{
		ConfigInfoUUID: c.UUID,
		Name:           c.Name,
		File:           c.File,
	}
	return cf.Add()

}

func (c *RepoConfig) Load() error {
	return nil
}

func (c *RepoConfig) Apply(uuid string) error {

	return nil
}
