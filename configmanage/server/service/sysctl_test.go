/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Thu Nov 7 10:19:32 2024 +0800
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

func TestSysctlConfig_Load(t *testing.T) {
	// 设置测试数据
	sysc := &SysctlConfig{
		ConfigInfoUUID: "83a15f95-430c-4889-aa60-b27624a81703",
	}
	err := sysc.Load()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("sysc: %v\n", sysc)
}

func TestGetSysctlFileByUUID(t *testing.T) {
	uuid := "9eb2dda0-1005-4bfa-acab-3daa9b6fbcc6"
	sysf, err := internal.GetSysctlFileByUUID(uuid)
	if err != nil {
		fmt.Printf("get sysctl file error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("SysctlFile: %v\n", sysf)
}

func TestGetSysctlFilesByConfigUUID(t *testing.T) {
	// 设置测试数据
	sysdcUUID := "83a15f95-430c-4889-aa60-b27624a81703"

	// 调用被测试的函数
	files, err := GetSysctlFilesByConfigUUID(sysdcUUID)
	if err != nil {
		fmt.Printf("load sysctl files error: %s\n", err)
		os.Exit(-1)
	}
	if len(files) == 0 {
		fmt.Printf("files is empty: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(files))
}

func TestGetSysctlFilesByNode(t *testing.T) {
	// 设置测试数据
	nodeid := "55555555-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	syscs, err := GetSysctlFilesByNode(nodeid)
	if err != nil {
		fmt.Printf("GetSysctlFilesByNode error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(syscs))
}
