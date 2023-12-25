<template>
  <div class="">
    <el-table
      :data="tableData"
      style="width: 100%; height: 90%"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column prop="id" label="编号" width="80" />
      <el-table-column prop="task_name" label="名称" />
      <el-table-column prop="command" label="命令" />
      <el-table-column prop="command" label="模板">
        <template #default="props">
          <el-link type="primary" @click="atuneDetail"
            >{{ props.row.command }}
            <el-icon class="el-icon--right"><icon-view /></el-icon
          ></el-link>
        </template>
      </el-table-column>
      <el-table-column prop="task_status" label="状态">
        <template #default="props">
          <el-progress
            v-if="props.row.task_status === '执行中'"
            :percentage="100"
            :format="format"
            striped
            striped-flow
            :duration="4"
          />
          <el-progress
            v-else-if="props.row.task_status === '已完成'"
            :percentage="100"
            status="success"
            :duration="0"
          />
          <el-progress
            v-else
            :percentage="0"
            :format="format"
            status="exception"
          />
        </template>
      </el-table-column>
      <el-table-column prop="create_time" label="创建时间" />
      <el-table-column prop="update_time" label="更新时间" />
      <el-table-column label="操作" width="80">
        <template #default="{ row }">
          <el-button
            round
            type="primary"
            size="small"
            @click="handleDetail(row)"
            >详情</el-button
          >
        </template>
      </el-table-column>
    </el-table>
  </div>
  <div class="pagination">
    <el-pagination
      v-model:current-page="currentPage"
      v-model:page-size="pageSize"
      :page-sizes="[10, 20, 50]"
      :total="totalItems"
      layout="total, sizes, prev, pager, next, jumper"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    ></el-pagination>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import { View as IconView } from "@element-plus/icons-vue";
import { getTaskLists, searchTune } from "@/api/atune";
import { Task, Atune } from "@/types/atune";

const tableData = ref([] as Task[]);
const currentPage = ref(1);
const pageSize = ref(10);
const totalItems = ref(0);
const emit = defineEmits(["selectionChange", "taskDetail", "atuneDetail"]);

let props = defineProps({
  refreshData: {
    type: Boolean,
    default: false,
  },
  searchTuneName: {
    type: String,
    default: "",
  },
  searchTune: {
    type: Boolean,
    default: false,
  },
});

// 状态渲染函数
const format = (percentage: number) => (percentage === 0 ? "error" : "running");

// 获取tune模板列表
const getTuneListsData = async () => {
  getTaskLists({ page: currentPage, size: pageSize }).then((res) => {
    if (res.data.code === 200) {
      tableData.value = res.data.data;
      totalItems.value = res.data.total;
    }
  });
};

// 高级搜索tune模板列表
const searchTuneListsData = () => {
  searchTune({
    page: currentPage.value,
    size: pageSize.value,
    name: props.searchTuneName,
  })
    .then((res) => {
      tableData.value = res.data.data;
      totalItems.value = res.data.total;
    })
    .catch((error) => {
      console.error("Error search data:", error);
    });
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

// 查看模板详情
const atuneDetail = (row: Atune) => {
  emit("atuneDetail", row);
};

// 查看任务详情
const handleDetail = (row: Task) => {
  emit("taskDetail", row);
};

// 选中多选框
const handleSelectionChange = (rows: Task[]) => {
  emit("selectionChange", rows);
};

onMounted(() => {
  getTuneListsData();
});

watch(
  () => [props.refreshData, props.searchTune],
  ([refreshData, searchTune]): void => {
    if (searchTune !== undefined) {
      searchTuneListsData();
    } else if (refreshData !== undefined) {
      getTuneListsData();
    }
  }
);
</script>

<style lang="less" scoped>
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
