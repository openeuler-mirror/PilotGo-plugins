package model

type Tunes struct {
	ID            int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TuneName      string `json:"tuneName"`
	CustomName    string `json:"custom_name"`
	Description   string `json:"description"`
	CreateTime    string `json:"create_time"`
	UpdateTime    string `json:"update_time"`
	WorkDirectory string `json:"workDir"`
	Prepare       string `json:"prepare"`
	Tune          string `json:"tune"`
	Restore       string `json:"restore"`
	Notes         string `json:"note"`
}

type RunResult struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TaskID      int    `json:"task_id"`
	MachineUUID string `json:"machine_uuid"`
	MachineIP   string `json:"machine_ip"`
	Command     string `json:"command"`
	RetCode     int    `json:"retcode"`
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
	IsUpdate    bool   `json:"is_update"`
}

type Tasks struct {
	ID         int         `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TuneID     int         `json:"tune_id"`
	TaskName   string      `json:"task_name"`
	Script     string      `json:"command"`
	TaskStatus string      `json:"task_status"`
	CreateTime string      `json:"create_time"`
	UpdateTime string      `json:"update_time"`
	RunResults []RunResult `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;" json:"results"`
	Tune       Tunes       `gorm:"foreignKey:TuneID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"tune"`
}

type AtuneClient struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	MachineUUID string `json:"machine_uuid"`
	MachineIP   string `json:"machine_ip"`
}
