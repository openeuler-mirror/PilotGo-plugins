package service

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/cmd/config/options"
	dangerousRule "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
)

type MySQLService struct {
	Conf *options.MysqlDBInfo
}

func (m *MySQLService) Name() string {
	return "MySQL"
}

func (m *MySQLService) Init(ctx *AppContext) error {
	err := ensureDatabase(m.Conf)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		m.Conf.UserName,
		m.Conf.Password,
		m.Conf.HostName,
		m.Conf.Port,
		m.Conf.DataBase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	db.AutoMigrate(&dangerousRule.DangerousRule{})

	ctx.MySQL = db
	return nil
}

func (m *MySQLService) Close() error {
	if App.MySQL == nil {
		return nil
	}
	sqlDB, err := App.MySQL.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func ensureDatabase(conf *options.MysqlDBInfo) error {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/?charset=utf8mb4&parseTime=true",
		conf.UserName,
		conf.Password,
		conf.HostName,
		conf.Port)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	createDBSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci",
		conf.DataBase,
	)
	if err := db.Exec(createDBSQL).Error; err != nil {
		return fmt.Errorf("create database error: %w", err)
	}

	d, err := db.DB()
	if err != nil {
		return err
	}
	defer d.Close()

	return nil
}
