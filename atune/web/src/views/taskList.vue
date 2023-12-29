<template>
  <div class="taskList">
    <my-table
      ref="taskRef"
      :get-data="getTaskLists"
      :get-all-data="getTaskLists"
      :del-func="deleteTask"
      :search-func="searchTask"
    >
      <template #listName>任务列表</template>
      <template #button_bar>
        <my-button @click="handleCreat">新增</my-button>
        <my-button @click="handleDelete">删除</my-button>
      </template>
      <el-table-column type="selection" width="55" />
      <el-table-column prop="id" label="编号" width="80" />
      <el-table-column prop="task_name" label="名称" />
      <el-table-column prop="command" label="命令" />
      <el-table-column label="模板编号">
        <template #default="props">
          <el-link type="primary" @click="atuneDetail(props.row)"
            >{{ props.row.tune ? props.row.tune_id : "暂无" }}
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
          <my-button round size="small" @click="handleDetail(row)"
            >详情</my-button
          >
        </template>
      </el-table-column>
    </my-table>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { View as IconView } from "@element-plus/icons-vue";
import { getTaskLists, searchTask, deleteTask } from "@/api/atune";
import { Task } from "@/types/atune";
import { useRouter } from "vue-router";
import { useAtuneStore } from "@/store/atune";

const taskRef = ref();
// 路由管理器
const router = useRouter();
const emit = defineEmits(["atuneDetail"]);

// 状态渲染函数
const format = (percentage: number) => (percentage === 0 ? "error" : "running");

// 查看模板详情
const atuneDetail = (row: Task) => {
  emit("atuneDetail", row);
};
// 查看任务详情
const handleDetail = (row: Task) => {
  router.push("/task/detail");
  useAtuneStore().setTaskRow(row);
};

// 新增
const handleCreat = () => {};
// 删除
const handleDelete = () => {
  taskRef.value.handleDelete();
};
</script>

<style lang="less" scoped>
.taskList {
  width: 100%;
  height: 100%;
}
</style>
