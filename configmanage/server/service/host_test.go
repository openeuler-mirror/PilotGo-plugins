/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Mon Nov 4 11:31:01 2024 +0800
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
		fmt.Printf("get hostfile error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("hc: %v\n", hf)
}

func TestGetHostFilesByConfigUUID(t *testing.T) {
	// 设置测试数据
	testUUID := "158e0acf-159b-4876-83b1-fa5f3d6460b1"

	// 调用被测试的函数
	files, err := GetHostFilesByConfigUUID(testUUID)
	if err != nil {
		fmt.Printf("load hostfiles error: %s\n", err)
		os.Exit(-1)
	}
	if len(files) == 0 {
		fmt.Printf("files is empty: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(files))
}

func TestGetHostFilesByNode(t *testing.T) {
	// 设置测试数据
	nodeid := "11111111-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	rcs, err := GetHostFilesByNode(nodeid)
	if err != nil {
		fmt.Printf("GetHostFilesByNode error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(rcs))
}
