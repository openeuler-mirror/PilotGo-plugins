/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Fri Nov 1 11:18:17 2024 +0800
 */
package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestRepoConfig_Record(t *testing.T) {
	// 设置测试数据
	rc := &RepoConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "9c3f8e3d-5f8e-42df-b2d0-49bf55cfeb56",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/root",
		Name:           "repo.txt",
		IsActive:       false,
	}

	// 调用被测试的函数
	err := rc.Record()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
}

func TestGetRopeFilesByCinfigUUID(t *testing.T) {
	// 设置测试数据
	testUUID := "9c3f8e3d-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	files, err := GetRopeFilesByConfigUUID(testUUID)
	if err != nil {
		fmt.Printf("load repofiles error: %s\n", err)
		os.Exit(-1)
	}
	if len(files) == 0 {
		fmt.Printf("files is empty: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(files))
}

func TestGetRopeFilesByNode(t *testing.T) {
	// 设置测试数据
	nodeid := "11111111-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	rcs, err := GetRopeFilesByNode(nodeid)
	if err != nil {
		fmt.Printf("GetRopeFilesByNode error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(rcs))
}
