<template>
  <div class="app">
    <div class="top">
      <span class="top-title">A-Tune调优管理</span>
      <span class="top-description">此处应有一段介绍</span>
    </div>
    <div class="nav">
      <!-- 两页 任务列表、调优模板列表 -->
      <router-link class="nav-link" to="/task" :active-class="'active'">
        任务</router-link
      >
      <router-link class="nav-link" to="/atune" :active-class="'active'">
        模板</router-link
      >
    </div>
    <el-config-provider :locale="zhCn">
      <router-view></router-view>
    </el-config-provider>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import locale from "element-plus/es/locale/lang/zh-cn";
const zhCn = ref(locale);
</script>

<style lang="less" scoped>
.app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100%;
  width: 100%;
  background-color: #f5f5f9;

  .top {
    width: 100%;
    height: 64px;
    display: flex;
    flex-direction: column;
    justify-content: center;

    &-title {
      padding-left: 40px;
      font-size: 18px;
      font-weight: 600;
      color: #333;
      display: inline-block;
    }

    &-description {
      padding-left: 40px;
      font-size: 16px;
      color: #444;
      font-weight: 500;
    }
  }
  .nav {
    width: 200px;
    height: 38px;
    margin: 0 auto;
    display: flex;
    justify-content: space-around;
    align-items: center;
    &-link {
      display: inline-block;
      width: 46%;
      text-align: center;
      text-decoration: none;
      color: rgb(96, 122, 207);
      border-radius: 6px;
      font-weight: bold;
      &:hover {
        background-color: rgba(96, 122, 207, 0.4);
        color: #fff;
      }
    }
    .active {
      color: #fff;
      background-color: rgb(96, 122, 207);
      box-shadow: 0 1px 6px 0 rgba(96, 122, 207, 0.4),
        0 2px 6px 0 rgba(96, 122, 207, 0.8);
    }
  }
}
</style>
