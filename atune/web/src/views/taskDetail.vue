<template>
  <div class="nav">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/task' }"
        >执行任务列表</el-breadcrumb-item
      >
      <el-breadcrumb-item>任务详情</el-breadcrumb-item>
    </el-breadcrumb>
  </div>
  <div class="task_detail shadow"></div>
</template>
<script lang="ts" setup>
import { useAtuneStore } from "@/store/atune";
import { ref } from "vue";
const { taskRow } = useAtuneStore();

// 获取路由参数
// console.log("详情数据：", taskRow);
</script>
<style lang="less" scope>
.nav {
  width: 96%;
  margin: 0 auto;
  height: 20px;
  text-align: left;
}
.task_detail {
  width: 98%;
  height: 100%;
  margin: 0 auto;
  background-color: #fff;
  border-radius: 6px;
}
</style>
