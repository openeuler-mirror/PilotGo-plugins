/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Wed Nov 6 09:40:00 2024 +0800
 */
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

func TestGetSSHDFilesByNode(t *testing.T) {
	// 设置测试数据
	nodeid := "44444444-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	sdcs, err := GetSSHDFilesByNode(nodeid)
	if err != nil {
		fmt.Printf("GetSSHDFilesByNode error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(sdcs))
}

func TestGetSSHDFilesByConfigUUID(t *testing.T) {
	// 设置测试数据
	sdcUUID := "5973e993-6236-4b53-9eb6-0cc23c652460"

	// 调用被测试的函数
	files, err := GetSSHDFilesByConfigUUID(sdcUUID)
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
