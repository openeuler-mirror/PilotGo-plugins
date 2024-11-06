package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
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

func TestSSHConfig_Load(t *testing.T) {
	// 设置测试数据
	sc := &SSHConfig{
		ConfigInfoUUID: "5973e993-6236-4b53-9eb6-0cc23c652460",
	}
	err := sc.Load()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("sshconfig: %v\n", sc)
}

func TestGetSSHFileByUUID(t *testing.T) {
	uuid := "64636f68-b6fa-426d-a253-ff03c6d529be"
	sf, err := internal.GetSSHFileByUUID(uuid)
	if err != nil {
		fmt.Printf("get ssh file error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("sshfile: %v\n", sf)
}

func TestGetSSHFilesByCinfigUUID(t *testing.T) {
	// 设置测试数据
	scUUID := "5973e993-6236-4b53-9eb6-0cc23c652460"

	// 调用被测试的函数
	files, err := GetSSHFilesByCinfigUUID(scUUID)
	if err != nil {
		fmt.Printf("load sshfiles error: %s\n", err)
		os.Exit(-1)
	}
	if len(files) == 0 {
		fmt.Printf("files is empty: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(files))
}

func TestGetSSHFilesByNode(t *testing.T) {
	// 设置测试数据
	nodeid := "33333333-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	scs, err := GetSSHFilesByNode(nodeid)
	if err != nil {
		fmt.Printf("GetSSHFilesByNode error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(scs))
}

func TestGetSSHDFilesByCinfigUUID(t *testing.T) {
	// 设置测试数据
	sdcUUID := "5973e993-6236-4b53-9eb6-0cc23c652460"

	// 调用被测试的函数
	files, err := GetSSHDFilesByCinfigUUID(sdcUUID)
	if err != nil {
		fmt.Printf("load sshdfiles error: %s\n", err)
		os.Exit(-1)
	}
	if len(files) == 0 {
		fmt.Printf("files is empty: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(files))
}
