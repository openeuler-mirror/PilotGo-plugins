<template>
  <div class="table">
    <el-table :data="tableData" style="width: 100%" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />
      <el-table-column prop="tuneName" label="名称" width="180" />
      <el-table-column prop="prepare" label="环境准备" />
      <el-table-column prop="tune" label="调优" />
      <el-table-column prop="restore" label="环境恢复" width="180" />
    </el-table>
  </div>
  <div class="pagination">
    <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[10, 20, 50]"
      :total="totalItems" layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
      @current-change="handleCurrentChange"></el-pagination>
  </div>
</template>

<script lang='ts' setup>
import { onMounted, ref, defineEmits, watch } from 'vue';
import { getTuneLists } from '@/api/atune'

const tableData = ref([] as Atune[]);
const currentPage = ref(1);
const pageSize = ref(10);
const totalItems = ref(0);
const emit = defineEmits(['selectionChange']);

export interface Atune {
  id: number
  tuneName: string
  workDir: string
  prepare: string
  tune: string
  restore: string
  note: string
}
let props = defineProps({
  refreshData: {
    type: Boolean,
    default: false
  }
})

// 获取tune模板列表
const fetchData = async () => {
  try {
    const response = await getTuneLists({
      page: currentPage.value,
      size: pageSize.value,
    });

    tableData.value = response.data.data;
    totalItems.value = response.data.total;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};

// 页码大小
const handleSizeChange = (newSize: number) => {
  pageSize.value = newSize;
  fetchData();
};

// 页数
const handleCurrentChange = (newPage: number) => {
  currentPage.value = newPage;
  fetchData();
};

// 选中多选框
const handleSelectionChange = (rows: Atune[]) => {
  emit('selectionChange', rows);
};

onMounted(() => {
  fetchData();
});

watch(() => props.refreshData, () => {
  fetchData();
});

</script>

<style lang = 'less' scoped>
.table {

  .el-table th,
  .el-table td {
    text-align: center; // 内容居中
  }

  .el-table th .cell {
    font-weight: bold; // 设置标签为黑体
  }
}

.pagination {
  position: absolute;
  bottom: 0;
  padding-bottom: 3%;
}
</style>
