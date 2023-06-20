package database

type Agent struct {
	ID             uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UUID           string `gorm:"not null" json:"uuid"`
	IP             string `gorm:"not null" json:"ip"`
	Port           string `gorm:"not null" json:"port"`
	Department     string `json:"department"`
	Statu          string `gorm:"not null" json:"state"`
	Gopher_deploy  bool   // ture:installed false:uninstalled
	Gopher_running bool   // true:running false:not running
	Gopher_version string `gorm:"type:varchar(20)"`
}
