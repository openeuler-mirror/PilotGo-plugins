package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/global"
)

func TestConfigInstance_Add(t *testing.T) {
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
