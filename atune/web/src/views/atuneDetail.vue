<template>
  <div class="atuneDetail">
    <el-container style="width: 100%; height: 100%">
      <el-header height="20%">
        <el-descriptions title="User Info">
          <el-descriptions-item label="Username"
            >kooriookami</el-descriptions-item
          >
          <el-descriptions-item label="Telephone"
            >18100000000</el-descriptions-item
          >
          <el-descriptions-item label="Place">Suzhou</el-descriptions-item>
          <el-descriptions-item label="Remarks">
            <el-tag size="small">School</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="Address"
            >No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu
            Province</el-descriptions-item
          >
        </el-descriptions>
      </el-header>
      <el-container>
        <el-aside
          width="50%"
          style="
            display: flex;
            justify-content: center;
            background-color: aquamarine;
          "
          ><flow-chart
            style="
              width: 90%;
              height: 100%;
              margin: 0 auto;
              background-color: blueviolet;
            "
          ></flow-chart
        ></el-aside>
        <el-main>
          <div class="shell">{{ "qweqweqeqr/qweqeqwew" }}</div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script lang="ts" setup>
import flowChart from "@/components/flowChart.vue";
import { ref } from "vue";
</script>
<style lang="less" scoped>
.atuneDetail {
  width: 100%;
  height: 100%;
  .shell {
    width: 90%;
    height: 100%;
    text-align: center;
  }
}
</style>
