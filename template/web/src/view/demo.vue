<template>
  <div>
    <p>
      一个简单的demo页面
    </p>
    <p>无参数请求接口结果：</p>
    <code>
      {{ result }}
    </code>
    <p>带参数请求接口结果：</p>
    <code>
      {{ result_params }}
    </code>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { demoAPI } from '../api/demo';

let result = ref('接口结果展示');
let result_params = ref('带参数接口结果展示');

demoAPI().then((res: any) => {
  if (res.data.code === 200) {
    result.value = res.data.data;
  }
})
demoAPI({ page: 1, size: 10 }).then((res: any) => {
  if (res.data.code === 200) {
    result_params.value = res.data.data;
  }
})
</script>

<style scoped></style>