<template>
  <div class="top">
    <span class="top-title">A-Tune调优管理</span>
  </div>
  <el-tree
    :data="atuneTree"
    :props="defaultProps"
    :default-checked-keys="selectedKeys"
    :highlight-current="true"
    @node-click="handleNodeClick"
  ></el-tree>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { getAtuneAllName } from '@/api/atune';

const atuneTree = ref([]);
const selectedKeys = ref<string[]>([]);
const defaultProps = ref({
  label: 'label',
});

onMounted(() => {
  getAtuneAllName();
});

getAtuneAllName().then((res) => {
  console.log(res);
  atuneTree.value = res.data.data.map((item: string, index: number) => ({
    label: item,
    key: index.toString(),
  }));
});
getAtuneAllName();

function handleNodeClick(node: { key: string }) {
 // 处理节点点击事件
 if (!selectedKeys.value.includes(node.key)) {
    // 如果节点未选中，将其 key 添加到 selectedKeys 中
    selectedKeys.value.push(node.key);
  } else {
    // 如果节点已选中，将其 key 从 selectedKeys 中移除
    selectedKeys.value = selectedKeys.value.filter((key) => key !== node.key);
  }
}
</script>

<style lang="less" scoped>
.top {
  width: 97.4%;
  margin: 0 auto;
  height: 64px;
  display: flex;
  justify-content: space-between;
  align-items: center;

  &-title {
    font-size: 20px;
    color: #222;
    font-weight: 500;
    display: inline-block;
  }
  .el-tree-node.is-checked {
    background-color: #9755a4 ; /* 自定义选中高亮的背景颜色 */
    color: #b7bc53; /* 设置文字颜色为白色 */
  }
}

</style>
