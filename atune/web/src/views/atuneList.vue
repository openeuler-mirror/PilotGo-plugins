<template>
  <div class="tuneList shadow" v-if="!showDetail">
    <my-table
      ref="tuneRef"
      :get-data="getTuneLists"
      :get-all-data="getTuneLists"
      :del-func="deleteTune"
      :search-func="searchTune"
    >
      <template #listName>模板列表</template>
      <template #button_bar>
        <my-button @click="handleCreat">新增</my-button>
        <my-button @click="handleDelete">删除</my-button>
      </template>
      <el-table-column type="selection" width="55" />
      <el-table-column prop="id" label="编号" width="100" />
      <el-table-column prop="tuneName" label="调优对象" width="100" />
      <el-table-column prop="custom_name" label="自定义名称" width="100" />
      <el-table-column prop="description" label="概述" />
      <el-table-column prop="create_time" label="创建时间" />
      <el-table-column prop="update_time" label="更新时间" />
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <my-button size="small" @click="handleDetail(row)">详情</my-button>
          <my-button size="small" @click="handleEdit(row)">编辑</my-button>
        </template>
      </el-table-column>
    </my-table>
  </div>
  <el-dialog title="调优模板信息" width="70%" v-model="showDialog">
    <atuneTemplete
      :is-tune="true"
      :selectedNodeData="selectedNodeData"
      :selectedEditRow="selectedEditRow"
      @closeDialog="closeDialog"
      @dataUpdated="handleDataUpdated"
    >
    </atuneTemplete>
  </el-dialog>
  <div class="tuneList shadow" v-if="showDetail">
    <router-view />
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { ElDialog } from "element-plus";
import { getTuneLists, searchTune, deleteTune } from "@/api/atune";
import atuneTemplete from "@/components/atuneTemplete.vue";
import { Atune } from "@/types/atune";
import { useRouter } from "vue-router";
import { useAtuneStore } from "@/store/atune";
const tuneRef = ref();
const router = useRouter();
const showDetail = ref(false);
const selectedNodeData = ref("");
const showDialog = ref(false);
const selectedEditRow = ref();
// 关闭dialog弹框
const closeDialog = () => {
  showDialog.value = false;
};
// 新增
const handleCreat = () => {
  showDialog.value = true;
};
// 删除
const handleDelete = () => {
  tuneRef.value.handleDelete();
};
// 详情
const handleDetail = (row: Atune) => {
  useAtuneStore().setTuneRow(row);
  showDetail.value = true;
  router.push("/atune/detail");
};
// 编辑
const handleEdit = (row: Atune) => {
  selectedEditRow.value = row;
  showDialog.value = true;
};
// 刷新
const handleDataUpdated = () => {
  tuneRef.value.handleRefresh();
};
</script>

<style lang="less" scoped>
.tuneList {
  width: 98%;
  margin: 0 auto;
  height: calc(100% - 44px - 10px);
}
</style>
