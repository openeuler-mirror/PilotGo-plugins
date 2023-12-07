<template>
  <div class="table-container">
    <div class="table-wrapper">
      <el-table :data="tableData" style="width: 100%;height:90%" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column type="expand" width="30">
          <template #default="props">
            <div class="expand-cell">
              <p class="note-title">注意事项:</p>
              <div class="note-content">{{ props.row.note }}</div>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="tuneName" label="名称" width="150" />
        <el-table-column prop="prepare" label="环境准备" width="370" />
        <el-table-column prop="tune" label="调优" width="400" />
        <el-table-column prop="restore" label="环境恢复" width="370" />
        <el-table-column label="操作" width="80">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="pagination">
      <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[10, 20, 50]"
        :total="totalItems" layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
        @current-change="handleCurrentChange"></el-pagination>
    </div>
  </div>
</template>

<script lang='ts' setup>
import { onMounted, ref, watch } from 'vue';
import { getTuneLists, searchTune } from '@/api/atune'
import { Atune } from '@/types/atune'

const tableData = ref([] as Atune[]);
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
const handleEdit = (row: Atune) => {
  emit('selectionEdit', row);
}

// 选中多选框
const handleSelectionChange = (rows: Atune[]) => {
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
.table-container {
  position: relative;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
}

.table-wrapper {
  overflow: auto;
  height: 82vh; 
  margin-left: 10px;
}

.pagination {
  position: absolute;
  bottom: 0;
  padding-bottom: 10px;
  padding-left: 20px;
  z-index: 1;
}

.expand-cell {
  display: flex;
  flex-direction: column;
  width: 96%;
  margin-left: auto;
}

.note-title {
  word-wrap: break-word; /* 换行 */
}

.note-content {
  word-wrap: break-word; 
  line-height: 1.8;
  padding-left: 2em;
  white-space: pre-line;
}
</style>
