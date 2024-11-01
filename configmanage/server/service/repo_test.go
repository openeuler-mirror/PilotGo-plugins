package service

import (
	"fmt"
	"os"
	"testing"

	"openeuler.org/PilotGo/configmanage-plugin/config"
	"openeuler.org/PilotGo/configmanage-plugin/db"
)

func TestGetRopeFilesByCinfigUUID(t *testing.T) {
	// 设置测试数据
	testUUID := "test-uuid"

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
	err := config.Init("")
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
