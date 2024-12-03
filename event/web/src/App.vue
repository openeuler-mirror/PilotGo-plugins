<!--
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: 赵振芳 <zhaozhenfang@kylinos.cn>
 * Date: Mon Nov 25 17:07:09 2024 +0800
-->
<template>
  <el-config-provider :locale="zhCn">
  <!-- 
    1.input 筛选类型，模糊搜索 
    2.time select 带快速选选择
    3.标题：跟pilotGo插件保持一致
  -->
  <my-table ref="tableRef" row-key="date" :getData="getData" style="width: 100%">
    <template #listName>事件列表</template>
      <template #button_bar>
        <el-button>btn</el-button>
      </template>
      <!-- <el-table-column type="selection" width="55" /> -->
      <el-table-column prop="id" label="编号" width="80" />
      <el-table-column prop="event_name" label="事件名称" />
      <el-table-column prop="create_time" label="创建时间" />
      <el-table-column prop="update_time" label="更新时间" />
      <el-table-column prop="description" label="描述" />
      <el-table-column label="操作" width="240">
        <template #default="{ row }">
          <el-button round size="small">btn</el-button>
        </template>
      </el-table-column>
  </my-table>
</el-config-provider>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import type { TableInstance } from 'element-plus'
import locale from 'element-plus/es/locale/lang/zh-cn';
import MyTable from './components/MyTable.vue';

const zhCn = ref(locale); 

const tableRef = ref<TableInstance>()

const getData = () => {

}
</script>