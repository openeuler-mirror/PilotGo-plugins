package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestSSHConfig_Record(t *testing.T) {
	// 设置测试数据
	sc := &SSHConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "5973e993-6236-4b53-9eb6-0cc23c652460",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/ssh",
		Name:           "ssh.txt",
		IsActive:       false,
	}

	// 调用被测试的函数
	err := sc.Record()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
}
