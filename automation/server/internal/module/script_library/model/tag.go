package model

type Tag struct {
	ID                  int    `json:"id" gorm:"primaryKey;autoIncrement;comment:标签ID"`
	Name                string `json:"name" gorm:"type:varchar(50);not null;unique;comment:标签名称"`
	Description         string `json:"description" gorm:"type:varchar(255);comment:标签描述"`
	Creator             string `json:"creator" gorm:"type:varchar(50);comment:创建者"`
	CreatedAt           string `json:"created_at" gorm:"comment:创建时间"`
	LastModifyUser      string `json:"last_modify_user" gorm:"type:varchar(50);comment:最后修改人"`
	LastModifyUpdatedAt string `json:"last_modify_updated_at" gorm:"comment:最后修改时间"`
}
