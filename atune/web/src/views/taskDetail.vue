<template>
  <div class="task_detail shadow"></div>
</template>
<script lang="ts" setup>
import { useAtuneStore } from "@/store/atune";
const { taskRow } = useAtuneStore();

// 获取路由参数
console.log("详情数据：", taskRow);
</script>
<style lang="less" scope>
.task_detail {
  width: 98%;
  height: 100%;
  margin: 0 auto;
  background-color: #fff;
  border-radius: 6px;
}
</style>
