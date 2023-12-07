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
      <div class="title">
        <div class="title-content">
          <p>调优模板</p>
        </div>
        <el-input v-model="searchTuneName" placeholder="请输入调优名称进行搜索..." :prefix-icon="Search" clearable
          style="width: 280px;margin-right: 10px;" @keydown.enter.native="handleSearch"></el-input>
        <el-button :icon="Search" @click="handleSearch">搜索</el-button>
        <el-button class="delete-button" @click="handleDelete">删除</el-button>
      </div>
      <div class="table">
        <atuneList @selectionChange="handleSelectionChange" @selectionEdit="handleSelectEdit" :refreshData="refreshData"
          :searchTuneName="searchTuneName" :searchTune="searchTune">
        </atuneList>
      </div>
    </div>
  </div>
  <el-dialog title="调优模板信息" width="70%" v-model="showDialog">
    <atuneTemplete :selectedNodeData="selectedNodeData" :selectedEditRow="selectedEditRow" @closeDialog="closeDialog"
      @dataUpdated="handleDataUpdated">
    </atuneTemplete>
  </el-dialog>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { getAtuneAllName, deleteTune } from '@/api/atune';
import { ElTree, ElDialog, ElMessage, ElMessageBox } from 'element-plus';
import { Search } from '@element-plus/icons-vue'
import atuneList from '@/components/atuneList.vue'
import atuneTemplete from '@/components/atuneTemplete.vue'
import { Atune } from '@/types/atune'

const atuneTree = ref([]);
const selectedNodeData = ref("")
const searchTuneName = ref("")
const searchTune = ref(false)
const showDialog = ref(false)
const selectedRows = ref([] as Atune[])
const selectedEditRow = ref()
const refreshData = ref(false)
const defaultProps = ref({
  label: 'label',
});

// 选中调优对象
const handleNodeClick = (node: any) => {
  selectedNodeData.value = node.label;
  showDialog.value = true;
}

// 关闭dialog弹框
const closeDialog = () => {
  showDialog.value = false;
}

// 选中多选框
const handleSelectionChange = (selected_Rows: any) => {
  selectedRows.value = selected_Rows;
}

// 编辑
const handleSelectEdit = (editRow: any) => {
  selectedEditRow.value = editRow
  showDialog.value = true;
}

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
  ElMessageBox.confirm('确定要删除吗？', '提示', {
    type: 'warning',
    confirmButtonText: '确定',
    cancelButtonText: '取消'
  })
    .then(() => {
      let ids = ref<number[]>([])
      selectedRows.value.forEach(item => {
        ids.value.push(item.id)
      })
      deleteTune({ ids: ids.value }).then(res => {
        if (res.data.code === 200) {
          ElMessage.success(res.data.msg)
          refreshData.value = !refreshData.value;
        } else {
          ElMessage.error(res.data.msg)
        }
      }).catch(err => {
        ElMessage.error("数据传输失败，请检查：", err)
      });
    })
}

onMounted(async () => {
  const res = await getAtuneAllName();
  atuneTree.value = res.data.data.map((item: string, index: number) => ({
    label: item,
    key: index.toString(),
  }));
});
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
    width: 99%;
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
