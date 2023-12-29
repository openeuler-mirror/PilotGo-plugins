<template>
  <div class="container">
    <!-- 任务列表 -->
    <div class="container-table shadow" v-show="!showDetail">
      <taskList @atuneDetail="handleAtuneDetail"> </taskList>
    </div>
    <router-view />
  </div>

  <!-- 调优模板详情 -->
  <el-dialog title="调优模板信息" width="70%" v-model="showDialog">
    <atuneTemplete
      :is-tune="false"
      :selectedEditRow="selectedEditRow"
      @closeDialog="closeDialog"
    >
    </atuneTemplete>
  </el-dialog>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { ElDialog } from "element-plus";
import taskList from "./taskList.vue";
import atuneTemplete from "@/components/atuneTemplete.vue";
import { Task, Atune } from "@/types/atune";
import { useRoute } from "vue-router";
import { onBeforeRouteUpdate } from "vue-router";
import { useRouterStore } from "@/store/router";

const showDialog = ref(false);
const selectedEditRow = ref();
const showDetail = ref(false);

// 每次刷新界面都需重新判断路由
onMounted(() => {
  showDetail.value = useRouterStore().showRoute(useRoute().fullPath, "detail");
});
// 组件内守卫
onBeforeRouteUpdate((to: any, _from: any, next: any) => {
  showDetail.value = useRouterStore().showRoute(to.fullPath, "detail");
  next();
});

// 关闭dialog弹框
const closeDialog = () => {
  showDialog.value = false;
};

// 查看模板详情
const handleAtuneDetail = (taskRow: Task | Atune) => {
  selectedEditRow.value = taskRow.tune;
  showDialog.value = true;
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
}
</style>
