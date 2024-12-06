<!--
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: 赵振芳 <zhaozhenfang@kylinos.cn>
 * Date: Mon Nov 25 17:07:09 2024 +0800
-->
<template>
  <el-config-provider :locale="zhCn">
  <my-table ref="tableRef" row-key="date" :getData="getEventList" :params="params" style="width: 100%">
    <template #listName>事件列表</template>
    <template #button_bar>
      <div class="date_div">
        <span>时间选择：</span>
        <el-date-picker
          v-model="dateRange"
          type="daterange"
          :shortcuts="shortcuts"
          range-separator="-"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
          @change="handleDate"
        />
      </div>&emsp;
      <span>类型搜索：</span>
      <el-select
      v-model="type"
      placeholder="请选择类型"
      clearable
      style="width: 240px"
      @change="handleSelect"
    >
      <el-option
        v-for="item in types"
        :key="item"
        :label="item"
        :value="item"
      />
    </el-select>
      &emsp;
      <el-button @click="handleSearch">搜索</el-button>
    </template>
    <el-table-column label="编号" type="index" width="80"/>
    <el-table-column prop="msg_type" label="类型" width="160"/>
    <el-table-column  label="时间" width="240">
    <template #default={row}>
    {{ row.time.split('.')[0] }}
    </template>
    </el-table-column>
    <el-table-column prop="value" label="消息内容" />
  </my-table>
</el-config-provider>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import locale from 'element-plus/es/locale/lang/zh-cn';
import MyTable from './components/MyTable.vue';

import { getEventList } from './apis/event';
import { shortcuts } from './utils';

const zhCn = ref(locale); 
const tableRef = ref()
const dateRange = ref([new Date(new Date().setDate(new Date().getDate() - 1)),new Date()])
const type = ref('');
const types=ref([] as string[]);
onMounted(() => {
  // 设置类型搜索下拉框内容
  getEventList(params.value).then(res => {
    if(res.data.code === 200) {
      types.value = res.data.msgType;
    }
  })
})

const params = ref({
  start:new Date(new Date().setDate(new Date().getDate() - 1)),
  stop:new Date(),
  search:''
})
// 时间范围筛选
const handleDate = (value:any) => {
  let c_start = new Date(value[0]);
  let c_stop = new Date(value[1]);
  let c_stopPlusday1 = new Date(c_stop.setDate(c_stop.getDate() + 1));
  params.value.start = new Date(c_start);
  params.value.stop = new Date(c_stopPlusday1);
}

// 选择类型方法
const handleSelect = () => {
  params.value.search=type.value;
}

// 搜索
const handleSearch = () => {
  tableRef.value?.getTableData();
}
</script>