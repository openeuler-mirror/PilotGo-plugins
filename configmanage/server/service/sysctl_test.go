package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestSysctlConfig_Record(t *testing.T) {
	// 设置测试数据
	sysc := &SysctlConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "83a15f95-430c-4889-aa60-b27624a81703",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/sysctl",
		Name:           "sysctl.txt",
		IsActive:       false,
	}

	// 调用被测试的函数
	err := sysc.Record()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
}
