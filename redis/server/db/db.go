package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"openeuler.org/PilotGo/redis-plugin/config"
	"openeuler.org/PilotGo/redis-plugin/global"
)

var Url string

type MysqlManager struct {
	ip       string
	port     int
	username string
	password string
	dbname   string
	db       *gorm.DB
}

func MysqldbInit(conf *config.MysqlDBInfo) error {
	m := &MysqlManager{
		ip:       conf.HostName,
		port:     conf.Port,
		username: conf.UserName,
		password: conf.Password,
		dbname:   conf.DataBase,
	}
	Url = fmt.Sprint("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		m.username,
		m.password,
		m.ip,
		m.port,
		m.dbname)
	var err error
	m.db, err = gorm.Open(mysql.Open(Url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	global.GlobalDB = m.db

	var db *sql.DB
	if db, err = m.db.DB(); err != nil {
		return err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	global.GlobalDB.AutoMigrate()
	return nil
}
