<template>
  <div class="">
    <el-table :data="tableData" style="width: 100%;height:90%" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />
      <el-table-column prop="id" label="编号" />
      <el-table-column prop="name" label="名称" />
      <el-table-column prop="atune" label="调优模板" />
      <el-table-column prop="state" label="状态">
        <template #default="props">
          <!-- 成功1 失败0 执行中2 -->
          <div v-if="props.state == 1">成功</div>
          <div v-else-if="props.state == 0">失败</div>
          <div v-else>执行中</div>
        </template>
      </el-table-column>
      <el-table-column prop="createTime" label="创建时间" />
      <el-table-column prop="updateTime" label="更新时间" />
      <el-table-column label="操作" width="80">
        <template #default="{ row }">
          <el-button type="primary" size="small" @click="handleEdit(row)">详情</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
  <div class="pagination">
    <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[10, 20, 50]"
      :total="totalItems" layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
      @current-change="handleCurrentChange"></el-pagination>
  </div>
</template>

<script lang='ts' setup>
import { onMounted, ref, watch } from 'vue';
import { getTaskLists, searchTune } from '@/api/atune'
import { Task } from '@/types/atune'

const tableData = ref([] as Task[]);
const currentPage = ref(1);
const pageSize = ref(10);
const totalItems = ref(0);
const emit = defineEmits(['selectionChange', 'selectionEdit']);

let props = defineProps({
  refreshData: {
    type: Boolean,
    default: false
  },
  searchTuneName: {
    type: String,
    default: ""
  },
  searchTune: {
    type: Boolean,
    default: false
  },
})

// 获取tune模板列表
const getTuneListsData = async () => {
  // 虚拟数据
  tableData.value = [{
    id: 1,
    name: '任务XXX',
    atune: 'nginx',
    state: 0,
    createTime: '2023-12-09',
    updateTime: '2023-12-19',
  }]
  getTaskLists().then(res => {
    if (res.data.code === 200) {
      tableData.value = res.data.data;
      totalItems.value = res.data.total;
    }
  })
};

// 高级搜索tune模板列表
const searchTuneListsData = () => {
  searchTune({
    page: currentPage.value,
    size: pageSize.value,
    name: props.searchTuneName
  }).then(res => {
    tableData.value = res.data.data;
    totalItems.value = res.data.total;
  }).catch(error => {
    console.error('Error search data:', error);
  })

};

// 页码大小
const handleSizeChange = (newSize: number) => {
  pageSize.value = newSize;
  getTuneListsData();
};

// 页数
const handleCurrentChange = (newPage: number) => {
  currentPage.value = newPage;
  getTuneListsData();
};

// 编辑
const handleEdit = (row: Task) => {
  emit('selectionEdit', row);
}

// 选中多选框
const handleSelectionChange = (rows: Task[]) => {
  emit('selectionChange', rows);
};

onMounted(() => {
  getTuneListsData();
});

watch(() => [props.refreshData, props.searchTune], ([refreshData, searchTune]): void => {
  if (searchTune !== undefined) {
    searchTuneListsData();
  } else if (refreshData !== undefined) {
    getTuneListsData();
  }
});

</script>

<style lang = 'less' scoped>
.expand-cell {
  display: flex;
  flex-direction: column;
  width: 96%;
  margin-left: auto;
}

.note-title {
  word-wrap: break-word;
  /* 换行 */
}

.note-content {
  word-wrap: break-word;
  line-height: 1.8;
  padding-left: 2em;
  white-space: pre-line;
}

.pagination {
  position: absolute;
  bottom: 0;
  padding-bottom: 10px;
  padding-left: 20px;
  z-index: 1;
}
</style>
