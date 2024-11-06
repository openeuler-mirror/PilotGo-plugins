package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestSSHDConfig_Record(t *testing.T) {
	// 设置测试数据
	sdc := &SSHDConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "4d415c77-5a3d-45fb-a221-67dba74db56d",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/ssh",
		Name:           "ssh.txt",
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
	fmt.Printf("sdc: %v\n", sdc)
}
