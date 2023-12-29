<template>
  <div class="atuneDetail">
    <el-container class="atuneDetail_con">
      <el-header height="20%">
        <el-descriptions title="模板详情">
          <el-descriptions-item label="名称:">
            <el-tag>{{ tuneName }}</el-tag></el-descriptions-item
          >
          <el-descriptions-item label="工作目录:">{{
            workDir
          }}</el-descriptions-item>
          <el-descriptions-item label="注意事项:">{{
            note
          }}</el-descriptions-item>
        </el-descriptions>
      </el-header>
      <el-container>
        <el-aside width="40%" style="display: flex; justify-content: center"
          ><flow-chart
            style="width: 90%; height: 100%; margin: 0 auto"
            @clickRect="showShell"
          ></flow-chart
        ></el-aside>
        <el-main>
          <div class="shell">{{ shell }}</div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script lang="ts" setup>
import flowChart from "@/components/flowChart.vue";
import { ref } from "vue";
import { useAtuneStore } from "@/store/atune";
const shell = ref("start:点击查看每一步具体的脚本信息");
let { prepare, restore, tune, workDir, tuneName, note } =
  useAtuneStore().tuneRow;

const showShell = (shellName: string) => {
  switch (shellName) {
    case "prepare":
      shell.value = prepare;
      break;
    case "restore":
      shell.value = restore;
      break;
    case "tune":
      shell.value = tune;
      break;

    default:
      shell.value = "start:点击查看每一步具体的脚本信息";
      break;
  }
};
</script>
<style lang="less" scoped>
.atuneDetail {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  background-color: #fff;
  &_con {
    width: 100%;
    height: 96%;
  }
  .shell {
    width: 90%;
    height: 100%;
    text-align: center;
  }
}
</style>
