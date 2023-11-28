<template>
  <div class="top">
    <span class="top-title">A-Tune调优管理</span>
  </div>
  <div class="container">
    <div class="tree-container">
      <div class="title">可调优对象</div>
      <el-tree :data="atuneTree" :props="defaultProps" :highlight-current="true" @node-click="handleNodeClick"></el-tree>
    </div>
    <div class="table-container">
      <div class="title">调优模板</div>
      <div class="table">
        <atuneList></atuneList>
      </div>
    </div>
    <div>
      <el-dialog title="调优信息" width="50%" @close="closeDialog" v-model="showDialog">
        <atuneTemplete :selectedNodeData="selectedNodeData"></atuneTemplete>
      </el-dialog>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { getAtuneAllName } from '@/api/atune';
import { ElTree, ElDialog } from 'element-plus';
import atuneList from '@/components/atuneList.vue'
import atuneTemplete from '@/components/atuneTemplete.vue'

const atuneTree = ref([]);
const selectedNodeData = ref("");
const showDialog = ref(false);
const defaultProps = ref({
  label: 'label',
});

onMounted(async () => {
  const res = await getAtuneAllName();
  atuneTree.value = res.data.data.map((item: string, index: number) => ({
    label: item,
    key: index.toString(),
  }));
});

function handleNodeClick(node: any) {
  selectedNodeData.value = node.label;
  // 点击节点时显示dialog弹框
  showDialog.value = true;
}

// 关闭dialog弹框
function closeDialog() {
  showDialog.value = false;
}
</script>

<style lang="less" scoped>
.top {
  width: 100%;
  height: 64px;
  display: flex;
  justify-content: space-between;
  align-items: center;

  &-title {
    padding-left: 40px;
    font-size: 20px;
    color: #222;
    font-weight: 500;
    display: inline-block;
  }
}

.container {
  display: flex;
  min-width: 100%;
  height: 95%;
}

.tree-container,
.table-container {
  border-radius: 10px;
  overflow: hidden;
  margin: 10px;
  background-color: #fff;
}

.tree-container {
  height: 95%;
  width: 20%;
  border: 1px solid #ddd;
  margin-left: 30px;
}

.table-container {
  height: 95%;
  width: 76%;
  border: 2px solid #ddd;
  display: flex;
  flex-direction: column;
}

.table {
  flex: 1;
  padding: 10px;
}

.title {
  height: 30px;
  background-color: #395a9c;
  color: #fff;
  padding: 10px;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  text-indent: 15px;
}
</style>
