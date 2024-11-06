package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

func TestSSHDConfig_Record(t *testing.T) {
	// 设置测试数据
	sdc := &SSHDConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "4d415c77-5a3d-45fb-a221-67dba74db56d",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/ssh",
		Name:           "sshd.txt",
		IsActive:       false,
	}

	// 调用被测试的函数
	err := sdc.Record()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
}

func TestSSHDConfig_Load(t *testing.T) {
	// 设置测试数据
	sdc := &SSHDConfig{
		ConfigInfoUUID: "4d415c77-5a3d-45fb-a221-67dba74db56d",
	}
	err := sdc.Load()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("sshdconfig: %v\n", sdc)
}

func TestGetSSHDFileByUUID(t *testing.T) {
	uuid := "ab9b9ee6-8750-46ce-958c-4d3e26bbf877"
	sdf, err := internal.GetSSHDFileByUUID(uuid)
	if err != nil {
		fmt.Printf("get sshd file error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("sshdfile: %v\n", sdf)
}
