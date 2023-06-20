package service

import (
	"fmt"
	"testing"

	"gitee.com/openeuler/PilotGo-plugins/sdk/common"
)

func TestInstall(t *testing.T) {

	//var param *common.Batch
	param := &common.Batch{
		BatchUUID:     "001",
		DepartmentIDs: []string{"2", "15"},
		MachineUUIDs:  []string{"b6acfeec-375c-4856-5610-0cdgef8cdgf4 ", "336d63a9-ca45-4b44-8577-c3fa0c0a46e9"},
	}
	_, err := Install(param)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")
}
