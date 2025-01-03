/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Fri Nov 15 16:10:13 2024 +0800
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

func TestDNSConfig_Record(t *testing.T) {
	// 设置测试数据
	dc := &DNSConfig{
		UUID:           uuid.New().String(),
		ConfigInfoUUID: "cd1574a4-cdad-4a55-9561-9ef371456b90",
		Content:        json.RawMessage(`{"test": "test"}`),
		Path:           "/root",
		Name:           "resolv.conf",
		IsActive:       false,
	}

	// 调用被测试的函数
	err := dc.Record()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
}

func TestDNSConfig_Load(t *testing.T) {
	// 设置测试数据
	dc := &DNSConfig{
		ConfigInfoUUID: "cd1574a4-cdad-4a55-9561-9ef371456b90",
	}
	err := dc.Load()
	if err != nil {
		fmt.Printf("record error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("hc: %v\n", dc)
}

func TestGetDNSFileByUUID(t *testing.T) {
	uuid := "a134a449-d635-4f08-8dd8-1e3de2a6a509"
	df, err := internal.GetDNSFileByUUID(uuid)
	if err != nil {
		fmt.Printf("get dnsfile error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("hc: %v\n", df)
}

func TestGetDNSFilesByConfigUUID(t *testing.T) {
	// 设置测试数据
	testUUID := "cd1574a4-cdad-4a55-9561-9ef371456b90"

	// 调用被测试的函数
	files, err := GetDNSFilesByConfigUUID(testUUID)
	if err != nil {
		fmt.Printf("load dnsfile error: %s\n", err)
		os.Exit(-1)
	}
	if len(files) == 0 {
		fmt.Printf("files is empty: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(files))
}

func TestGetDNSFilesByNode(t *testing.T) {
	// 设置测试数据
	nodeid := "77777777-5f8e-42df-b2d0-49bf55cfeb56"

	// 调用被测试的函数
	dcs, err := GetDNSFilesByNode(nodeid)
	if err != nil {
		fmt.Printf("GetDNSFilesByNode error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(len(dcs))
}
