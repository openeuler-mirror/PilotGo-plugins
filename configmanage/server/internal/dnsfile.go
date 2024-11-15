package internal

import (
	"encoding/json"
	"time"

	"openeuler.org/PilotGo/configmanage-plugin/db"
)

type DNSFile struct {
	ID             int        `gorm:"autoIncrement:true"`
	UUID           string     `gorm:"primary_key;type:varchar(50)" json:"uuid"`
	ConfigInfo     ConfigInfo `gorm:"Foreignkey:ConfigInfoUUID"`
	ConfigInfoUUID string
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	Content        json.RawMessage `gorm:"type:json" json:"content"`
	Version        string          `gorm:"type:varchar(50)" json:"version"`
	IsActive       bool            `json:"isactive"`
	IsFromHost     bool            `gorm:"default:false" json:"isfromhost"`
	Hostuuid       string          `gorm:"type:varchar(50)" json:"hostuuid"`
	CreatedAt      time.Time
}

func (df *DNSFile) Add() error {
	sql := `
	INSERT INTO dns_file (uuid,config_info_uuid,path,name,content,version,is_active,is_from_host,hostuuid) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) 
	ON DUPLICATE KEY UPDATE
		uuid = VALUES(uuid),
		config_info_uuid = VALUES(config_info_uuid),
		path = VALUES(path),
		name = VALUES(name),
		content = VALUES(content),
		version = VALUES(version),
		is_active = VALUES(is_active),
		is_from_host = VALUES(is_from_host),
		hostuuid = VALUES(hostuuid);
	`
	return db.MySQL().Exec(sql,
		df.UUID,
		df.ConfigInfoUUID,
		df.Path,
		df.Name,
		df.Content,
		df.Version,
		df.IsActive,
		df.IsFromHost,
		df.Hostuuid,
	).Error
}
func GetDNSFileByInfoUUID(uuid string, isindex interface{}) (DNSFile, error) {
	var file DNSFile
	if isindex != nil {
		err := db.MySQL().Model(&DNSFile{}).Where("config_info_uuid=? && is_index = ?", uuid, isindex).Find(&file).Error
		return file, err
	}
	err := db.MySQL().Model(&DNSFile{}).Where("config_info_uuid=?", uuid).Find(&file).Error
	return file, err
}

func GetDNSFileByUUID(uuid string) (DNSFile, error) {
	var file DNSFile
	err := db.MySQL().Model(&DNSFile{}).Where("uuid=?", uuid).Find(&file).Error
	return file, err
}
