package model

type Tag struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement;comment:标签ID"`
	Name        string `json:"name" gorm:"type:varchar(50);not null;unique;comment:标签名称"`
	Description string `json:"description" gorm:"type:varchar(255);comment:标签描述"`
	ModifyUser  string `json:"modify_user" gorm:"type:varchar(50);comment:最后修改人"`
	ModifyTime  string `json:"modify_time" gorm:"comment:最后修改时间"`
}
