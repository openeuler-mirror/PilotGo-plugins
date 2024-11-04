package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

func TestHostConfig_Record(t *testing.T) {
	// 设置测试数据
	hc := &HostConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "158e0acf-159b-4876-83b1-fa5f3d6460b1",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/host",
		Name:           "host.txt",
		IsActive:       false,
	}

	// 调用被测试的函数
	err := hc.Record()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
}

func TestHostConfig_Load(t *testing.T) {
	// 设置测试数据
	hc := &HostConfig{
		ConfigInfoUUID: "158e0acf-159b-4876-83b1-fa5f3d6460b1",
	}
	err := hc.Load()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("hc: %v\n", hc)
}

func TestGetHostFileByUUID(t *testing.T) {
	uuid := "4254b485-8e8a-427c-bed1-5da05e363657"
	hf, err := internal.GetHostFileByUUID(uuid)
	if err != nil {
		fmt.Printf("get host file error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("hc: %v\n", hf)
}
