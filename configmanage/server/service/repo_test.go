package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/config"
	"openeuler.org/PilotGo/configmanage-plugin/db"
)

func TestRepoConfig_Record(t *testing.T) {
	// 设置测试数据
	rc := &RepoConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "9c3f8e3d-5f8e-42df-b2d0-49bf55cfeb56",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/root",
		Name:           "repo.txt",
		IsActive:       false,
	}

	// 调用被测试的函数
	err := rc.Record()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
}

func TestGetRopeFilesByCinfigUUID(t *testing.T) {
	// 设置测试数据
	testUUID := "9c3f8e3d-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	files, err := GetRopeFilesByCinfigUUID(testUUID)
	if err != nil {
		fmt.Printf("load repofiles error: %s\n", err)
		os.Exit(-1)
	}
	if len(files) == 0 {
		fmt.Printf("files is empty: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(files))
}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	err := config.Init(".././config.yaml")
	if err != nil {
		fmt.Printf("load config error: %s\n", err)
		os.Exit(-1)
	}
	err = db.MysqldbInit(config.Config().Mysql)
	if err != nil {
		fmt.Printf("init database error: %s\n", err)
		os.Exit(-1)
	}
	Init()
	m.Run()
	fmt.Println("end")
}
