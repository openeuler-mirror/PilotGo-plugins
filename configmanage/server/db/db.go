package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"openeuler.org/PilotGo/configmanage-plugin/config"
)

type MysqlManager struct {
	ip       string
	port     int
	userName string
	passWord string
	dbName   string
	db       *gorm.DB
}

var Url string
var global_db *gorm.DB

func MySQL() *gorm.DB {
	return global_db
}

func MysqldbInit(conf *config.MysqlDBInfo) error {
	m := &MysqlManager{
		ip:       conf.HostName,
		port:     conf.Port,
		userName: conf.UserName,
		passWord: conf.Password,
		dbName:   conf.DataBase,
	}
	err := ensureDatabase(m)
	if err != nil {
		return err
	}
	Url = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		m.userName,
		m.passWord,
		m.ip,
		m.port,
		m.dbName)
	m.db, err = gorm.Open(mysql.Open(Url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	global_db = m.db

	var db *sql.DB
	if db, err = m.db.DB(); err != nil {
		return err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	MySQL().AutoMigrate()
	return nil
}

func ensureDatabase(conf *MysqlManager) error {
	Url := fmt.Sprintf("%s:%s@(%s:%d)/?charset=utf8mb4&parseTime=true",
		conf.userName,
		conf.passWord,
		conf.ip,
		conf.port)
	db, err := gorm.Open(mysql.Open(Url))
	if err != nil {
		return err
	}

	creatDataBase := "CREATE DATABASE IF NOT EXISTS " + conf.ip + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci"
	db.Exec(creatDataBase)

	d, err := db.DB()
	if err != nil {
		return err
	}
	if err = d.Close(); err != nil {
		return err
	}
	return nil
}
