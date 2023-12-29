<template>
  <div class="container">
    <!-- 任务列表 -->
    <div class="container-table shadow" v-show="!showDetail">
      <taskList
        @taskDetail="handleTaskDetail"
        @atuneDetail="handleAtuneDetail"
        :refreshData="refreshData"
        :searchTuneName="searchTuneName"
        :searchTune="searchTune"
      >
      </taskList>
    </div>
    <!-- 单项任务详情 -->
    <div v-show="showDetail" class="container-nav">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/task' }" @click="returnHome"
          >执行任务列表</el-breadcrumb-item
        >
        <el-breadcrumb-item>任务详情</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <router-view />
  </div>

  <!-- 调优模板详情 -->
  <el-dialog title="调优模板信息" width="70%" v-model="showDialog">
    <atuneTemplete
      :is-tune="false"
      :selectedNodeData="selectedNodeData"
      :selectedEditRow="selectedEditRow"
      @closeDialog="closeDialog"
      @dataUpdated="handleDataUpdated"
    >
    </atuneTemplete>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { ElDialog } from "element-plus";
import taskList from "./taskList.vue";
import atuneTemplete from "@/components/atuneTemplete.vue";
import { useRouter } from "vue-router";
import { Task, Atune } from "@/types/atune";

const selectedNodeData = ref("");
const searchTuneName = ref("");
const searchTune = ref(false);
const showDialog = ref(false);
const selectedEditRow = ref();
const refreshData = ref(false);
const showDetail = ref(false);

// 路由管理器
const router = useRouter();

// 关闭dialog弹框
const closeDialog = () => {
  showDialog.value = false;
};

// 返回任务列表
const returnHome = () => {
  showDetail.value = false;
};

// 新增
const handleCreat = () => {};

// 查看模板详情
const handleAtuneDetail = (taskRow: Task | Atune) => {
  selectedEditRow.value = taskRow.tune;
  showDialog.value = true;
};

// 查看任务详情
const handleTaskDetail = () => {
  router.push("/task/detail");
  showDetail.value = true;
};

// 刷新
const handleDataUpdated = () => {
  refreshData.value = !refreshData.value;
};
</script>

<style lang="less" scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  height: calc(100% - 44px - 10px);

  &-table {
    width: 98%;
    margin: 0 auto;
    height: 100%;
  }

  &-nav {
    width: 96%;
    margin: 0 auto;
    height: 20px;
    text-align: left;
  }
}
</style>
