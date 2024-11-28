/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Tue Nov 26 15:36:42 2024 +0800
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

func TestPathConfig_Record(t *testing.T) {
	// 设置测试数据
	hc := &PathConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "158e0acf-159b-4876-83b1-fa5f3d6460b1",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/path",
		Name:           "path.txt",
		IsActive:       false,
	}

	// 调用被测试的函数
	err := hc.Record()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
}

func TestPathConfig_Load(t *testing.T) {
	// 设置测试数据
	pc := &PathConfig{
		ConfigInfoUUID: "158e0acf-159b-4876-83b1-fa5f3d6460b1",
	}
	err := pc.Load()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("pc: %v\n", pc)
}

func TestGetPathFileByUUID(t *testing.T) {
	uuid := "4254b485-8e8a-427c-bed1-5da05e363657"
	hf, err := internal.GetPathFileByUUID(uuid)
	if err != nil {
		fmt.Printf("get pathfile error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("hc: %v\n", hf)
}

func TestGetPathFilesByConfigUUID(t *testing.T) {
	// 设置测试数据
	testUUID := "158e0acf-159b-4876-83b1-fa5f3d6460b1"

	// 调用被测试的函数
	files, err := GetPathFilesByConfigUUID(testUUID)
	if err != nil {
		fmt.Printf("load PathConfigs error: %s\n", err)
		os.Exit(-1)
	}
	if len(files) == 0 {
		fmt.Printf("files is empty: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(files))
}

func TestGetPathFilesByNode(t *testing.T) {
	// 设置测试数据
	nodeid := "11111111-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	rcs, err := GetPathFilesByNode(nodeid)
	if err != nil {
		fmt.Printf("GetPathConfigsByNode error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(rcs))
}
