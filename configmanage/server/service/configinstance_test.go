package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/config"
	"openeuler.org/PilotGo/configmanage-plugin/db"
	"openeuler.org/PilotGo/configmanage-plugin/global"
)

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

func TestConfigInstanceTypeRepo_Add(t *testing.T) {
	ci := &ConfigInstance{
		UUID:        uuid.New().String(),
		Type:        global.Repo,
		Description: "test-repo-description",
		BatchIds:    []int{1, 2, 3},
		DepartIds:   []int{10, 20, 30},
		Nodes:       []string{"11111111-5f8e-42df-b2d0-49bf55cfeb56"},
	}
	err := ci.Add()
	if err != nil {
		fmt.Printf("Add() error = %v, want nil", err)
		os.Exit(-1)
	}
}

func TestGetInfoByUUID(t *testing.T) {
	testUUID := "9c3f8e3d-5f8e-42df-b2d0-49bf55cfeb56"

	configInfo, err := GetInfoByUUID(testUUID)
	if err != nil {
		fmt.Printf("GetInfoByUUID() error = %v, want nil", err)
	}
	fmt.Println(configInfo)
}
func TestGetConfigByUUID(t *testing.T) {
	testUUID := "9c3f8e3d-5f8e-42df-b2d0-49bf55cfeb56"

	configInfo, err := GetConfigByUUID(testUUID)
	if err != nil {
		fmt.Printf("GetConfigByUUID() error = %v, want nil", err)
	}
	fmt.Println(configInfo)
}

func TestGetInfos(t *testing.T) {
	total, data, err := GetInfos(0, 10)
	if err != nil {
		fmt.Printf("GetInfos() error = %v, want nil", err)
		os.Exit(-1)
	}
	fmt.Println(total, data)
}

func TestConfigInstanceTypeHost_Add(t *testing.T) {
	ci := &ConfigInstance{
		UUID:        uuid.New().String(),
		Type:        global.Host,
		Description: "test-host-description",
		BatchIds:    []int{4, 5},
		DepartIds:   []int{40, 50},
		Nodes:       []string{"22222222-5f8e-42df-b2d0-49bf55cfeb56"},
	}
	err := ci.Add()
	if err != nil {
		fmt.Printf("Add() error = %v, want nil", err)
		os.Exit(-1)
	}
}

func TestConfigInstanceTypeSSH_Add(t *testing.T) {
	ci := &ConfigInstance{
		UUID:        uuid.New().String(),
		Type:        global.SSH,
		Description: "test-SSH-description",
		BatchIds:    []int{6},
		DepartIds:   []int{60},
		Nodes:       []string{"33333333-5f8e-42df-b2d0-49bf55cfeb56"},
	}
	err := ci.Add()
	if err != nil {
		fmt.Printf("Add() error = %v, want nil", err)
		os.Exit(-1)
	}
}

func TestConfigInstanceTypeSSHD_Add(t *testing.T) {
	ci := &ConfigInstance{
		UUID:        uuid.New().String(),
		Type:        global.SSHD,
		Description: "test-SSHD-description",
		BatchIds:    []int{6},
		DepartIds:   []int{60},
		Nodes:       []string{"44444444-5f8e-42df-b2d0-49bf55cfeb56"},
	}
	err := ci.Add()
	if err != nil {
		fmt.Printf("Add() error = %v, want nil", err)
		os.Exit(-1)
	}
}
