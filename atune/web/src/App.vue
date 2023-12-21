<template>
  <div class="app">
    <el-config-provider :locale="zhCn">
      <router-view></router-view>
    </el-config-provider>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import locale from 'element-plus/es/locale/lang/zh-cn';
const zhCn = ref(locale);
</script>

<style lang="less" scoped>
.app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100%;
  width: 100%;
  // overflow: hidden; /* Ensure no scrollbars on the app container */
}

// html {
//   margin: 0;
//   padding: 0;
//   height: 100%;
//   width: 100%;
// }

// body {
//   margin: 0;
//   padding: 0;
//   height: 100%;
//   width: 100%;
//   background-color: rgb(11, 35, 117);
//   // overflow: hidden; /* Ensure no scrollbars on the body */
//   transform-origin: 0 0 0;
//   -moz-transform-origin: 0 0 0;
//   -ms-transform-origin: 0 0 0;
//   -webkit-transform-origin: 0 0 0;
//   -o-transform-origin: 0 0 0;
// }
</style>