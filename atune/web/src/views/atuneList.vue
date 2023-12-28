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
      <el-table-column prop="tuneName" label="名称" width="100" />
      <el-table-column prop="description" label="概述" />
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="更新时间" />
      <el-table-column label="操作" width="80">
        <template #default="{ row }">
          <my-button size="small" @click="handleDetail(row)">详情</my-button>
          <my-button size="small" @click="handleEdit(row)">编辑</my-button>
        </template>
      </el-table-column>
    </my-table>
  </div>
  <div class="tuneList shadow" v-if="showDetail">
    <router-view />
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { getTuneLists, searchTune, deleteTune } from "@/api/atune";
import { Atune } from "@/types/atune";
import { useRouter } from "vue-router";
const emit = defineEmits(["atuneDetail"]);
const tuneRef = ref();
const router = useRouter();
const showDetail = ref(false);
// 新增
const handleCreat = () => {};
// 删除
const handleDelete = () => {
  tuneRef.value.handleDelete();
};
// 详情
const handleDetail = (atuneRow: any) => {
  atuneRow;
  showDetail.value = true;
  router.push("/atune/detail");
};
// 编辑
const handleEdit = (row: Atune) => {
  emit("atuneDetail", row);
};
</script>

<style lang="less" scoped>
.tuneList {
  width: 98%;
  margin: 0 auto;
  height: calc(100% - 44px - 10px);
}
</style>
