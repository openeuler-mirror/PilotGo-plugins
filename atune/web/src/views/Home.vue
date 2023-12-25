<template>
  <div class="container">
    <div class="table-container shadow" v-show="!showDetail">
      <div class="title">
        <div class="title-content" style="font-size: 18px">执行任务列表</div>
        <el-input
          v-model="searchTuneName"
          placeholder="请输入任务名称进行搜索..."
          :prefix-icon="Search"
          clearable
          style="width: 280px; margin-right: 10px"
          @keydown.enter.native="handleSearch"
        ></el-input>
        <el-button :icon="Search" @click="handleSearch">搜索</el-button>
        <el-button class="delete-button" @click="handleCreat">新增</el-button>
        <el-button class="delete-button" @click="handleDelete">删除</el-button>
      </div>
      <!-- 任务列表 -->
      <div>
        <taskList
          @selectionChange="handleSelectionChange"
          @taskDetail="handleTaskDetail"
          @atuneDetail="handleAtuneDetail"
          :refreshData="refreshData"
          :searchTuneName="searchTuneName"
          :searchTune="searchTune"
        >
        </taskList>
      </div>
    </div>
    <!-- 单项任务详情 -->
    <div v-show="showDetail" class="container-nav">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/task' }" @click="returnHome"
          >执行任务列表</el-breadcrumb-item
        >
        <el-breadcrumb-item>任务详情</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <router-view />
  </div>

  <!-- 调优模板详情 -->
  <el-dialog title="调优模板信息" width="70%" v-model="showDialog">
    <atuneTemplete
      :selectedNodeData="selectedNodeData"
      :selectedEditRow="selectedEditRow"
      @closeDialog="closeDialog"
      @dataUpdated="handleDataUpdated"
    >
    </atuneTemplete>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { deleteTune } from "@/api/atune";
import { ElDialog, ElMessage, ElMessageBox } from "element-plus";
import { Search } from "@element-plus/icons-vue";
import taskList from "./taskList.vue";
import atuneTemplete from "@/components/atuneTemplete.vue";
import { Task } from "@/types/atune";
import router from "../router/index.ts";

const selectedNodeData = ref("");
const searchTuneName = ref("");
const searchTune = ref(false);
const showDialog = ref(false);
const selectedRows = ref([] as Task[]);
const selectedEditRow = ref();
const refreshData = ref(false);
const showDetail = ref(false);

// 关闭dialog弹框
const closeDialog = () => {
  showDialog.value = false;
};

// 选中多选框
const handleSelectionChange = (selected_Rows: any) => {
  selectedRows.value = selected_Rows;
};

// 返回任务列表
const returnHome = () => {
  showDetail.value = false;
};

// 新增
const handleCreat = () => {};

// 查看模板详情
const handleAtuneDetail = (editRow: any) => {
  selectedEditRow.value = editRow;
  showDialog.value = true;
};

// 查看任务详情
const handleTaskDetail = (editRow: any) => {
  router.push({
    path: "/task/detail",
    params: {
      row: editRow,
    },
  });
  showDetail.value = true;
};

// 刷新
const handleDataUpdated = () => {
  refreshData.value = !refreshData.value;
};

// 搜索
const handleSearch = () => {
  searchTune.value = !searchTune.value;
};

// 删除
const handleDelete = () => {
  ElMessageBox.confirm("确定要删除吗？", "提示", {
    type: "warning",
    confirmButtonText: "确定",
    cancelButtonText: "取消",
  }).then(() => {
    let ids = ref<number[]>([]);
    selectedRows.value.forEach((item) => {
      ids.value.push(item.id);
    });
    deleteTune({ ids: ids.value })
      .then((res) => {
        if (res.data.code === 200) {
          ElMessage.success(res.data.msg);
          refreshData.value = !refreshData.value;
        } else {
          ElMessage.error(res.data.msg);
        }
      })
      .catch((err) => {
        ElMessage.error("数据传输失败，请检查：", err);
      });
  });
};

/* // 获取所有的可调优对象
onMounted(async () => {
  const res = await getAtuneAllName();
  atuneTree.value = res.data.data.map((item: string, index: number) => ({
    label: item,
    key: index.toString(),
  }));
}); */
</script>

<style lang="less" scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-width: 100%;
  height: calc(100% - 64px - 50px);
  &-nav {
    width: 96%;
    margin: 0 auto;
    height: 20px;
    text-align: left;
  }

  .table-container {
    height: 95%;
    width: 98%;
    display: flex;
    flex-direction: column;
    border-radius: 12px;
  }

  .table {
    flex: 1;
    padding: 10px;
  }

  .title {
    height: 30px;
    background-color: rgba(96, 122, 207, 0.9);
    color: #fff;
    padding: 10px;
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;
    text-indent: 15px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .title-content {
      flex: 1;
    }

    .delete-button {
      margin-right: 10px;
    }
  }
}
</style>
