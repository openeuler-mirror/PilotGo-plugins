<template>
  <el-button class="button" v-bind="$attrs" plain>
    <slot></slot>
  </el-button>
</template>
<script lang="ts" setup></script>
<style lang="less" scoped>
.button {
  color: rgb(96, 122, 207) !important;
  &:hover {
    background-color: #fff !important;
    color: rgb(96, 122, 207) !important;
    border: 1px solid rgb(96, 122, 207) !important;
  }
}
</style>
