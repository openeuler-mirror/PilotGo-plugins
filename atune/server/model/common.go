package model

type Tunes struct {
	ID            int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TuneName      string `json:"tuneName"`
	WorkDirectory string `json:"workDir"`
	Prepare       string `json:"prepare"`
	Tune          string `json:"tune"`
	Restore       string `json:"restore"`
	Notes         string `json:"note"`
}

type RunResult struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
}
