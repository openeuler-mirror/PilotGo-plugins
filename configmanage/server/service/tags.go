/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Fri Nov 1 15:40:11 2024 +0800
 */
package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"openeuler.org/PilotGo/configmanage-plugin/global"
	"openeuler.org/PilotGo/configmanage-plugin/internal"
)

func GetTags() {
	tag_cb := func(uuids []string) []common.Tag {
		var tags []common.Tag
		for _, uuid := range uuids {
			ok := true
			cns, err := internal.GetConfigNodesByNode(uuid)
			if err != nil || len(cns) == 0 {
				ok = false
			}
			if ok {
				tag := common.Tag{
					UUID: uuid,
					Type: common.TypeOk,
					Data: "configmanage",
				}
				tags = append(tags, tag)
			} else {
				tag := common.Tag{
					UUID: uuid,
					Type: common.TypeError,
					Data: "",
				}
				tags = append(tags, tag)
			}
		}
		return tags
	}
	global.GlobalClient.OnGetTags(tag_cb)
}
