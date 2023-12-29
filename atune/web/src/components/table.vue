<template>
  <div class="my_table">
    <!-- table工具条 -->
    <el-row class="my_table_header">
      <div class="my_table_header_title">
        <slot name="listName"></slot>
      </div>
      <div class="my_table_header_operation">
        <!-- 模糊搜索 -->
        <div class="operation-select-input">
          <el-input
            v-model="keyWord"
            placeholder="请输入关键词进行搜索..."
            :prefix-icon="Search"
            clearable
            @keydown.enter.native="handleSearch"
            @clear="handleRefresh"
          ></el-input>
        </div>
      </div>
      <div class="my_table_header_button">
        <slot name="button_bar"></slot>
      </div>
    </el-row>
    <!-- 列表 -->
    <div class="my_table_content" ref="tableBox">
      <el-table
        ref="myTableRef"
        :data="tableData"
        class="table"
        @select="handleRowSelectionChange"
        @selection-change="handleSelectinChange"
        v-loading="loading"
      >
        <slot></slot>
        <template #append>
          <slot name="append"></slot>
        </template>
        <template #empty>
          <el-empty description="暂无数据"></el-empty>
        </template>
      </el-table>
    </div>
    <!-- 分页 -->
    <div class="my_table_page">
      <el-pagination
        v-model:current-page="page.currentPage"
        v-model:page-size="page.pageSize"
        popper-class="pagePopper"
        :page-sizes="[10, 20, 25, 50, 75, 100]"
        :small="page.small"
        :background="page.background"
        layout="total, sizes, prev, pager, next, jumper"
        :total="page.total"
        @size-change="getTableData"
        @current-change="getTableData"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { Search } from "@element-plus/icons-vue";
import { ElTable, ElMessage, ElMessageBox } from "element-plus";
import { ReaultData } from "@/types/atune";
const props = defineProps({
  getData: {
    type: Function,
    required: true,
  },
  getAllData: {
    type: Function,
    required: true,
  },
  searchFunc: {
    type: Function,
    required: true,
  },
  delFunc: {
    type: Function,
    required: false,
  },
});
const emit = defineEmits(["handleSelect", "handleRowclick"]);
const loading = ref(false);
const tableData = ref([] as any[]);
const all_datas = ref([]);
const myTableRef = ref<InstanceType<typeof ElTable>>();
const currentNum = ref(0); //复选框当前页数量
const keyWord = ref("");
const selectedRows = ref([] as any[]);
const page = reactive({
  total: 0,
  currentPage: 1,
  pageSize: 10,
  small: true,
  background: true,
});
onMounted(async () => {
  await getTableData();
  // getAllData();
});

// 获取表格数据
const getTableData = () => {
  loading.value = true;
  props.getData!({ page: page.currentPage, size: page.pageSize }).then(
    (res: { data: ReaultData }) => {
      let result: ReaultData = res.data;
      if (result && result.code === 200) {
        loading.value = false;
        tableData.value = result.data;
        currentNum.value = result.data.length;
        page.total = result.total;
      } else {
        loading.value = false;
        tableData.value = [];
        currentNum.value = 0;
        page.total = 10;
      }
    }
  );
};

// 搜索函数
const handleSearch = () => {
  props
    .searchFunc({
      page: page.currentPage,
      size: page.pageSize,
      search: keyWord.value,
    })
    .then((res: any) => {
      loading.value = false;
      tableData.value = res.data.data;
      currentNum.value = res.data.data.length;
      page.total = res.data.total;
    })
    .catch((error: any) => {
      console.error("Error search data:", error);
    });
};

// 获取全部不分页数据
const getAllData = () => {
  all_datas.value = [];
  props.getAllData({ paged: "false" }).then((res: any) => {
    if (res.data && res.data.code === 200) {
      all_datas.value = res.data.data;
    }
  });
};

// 表格被选择数据发生变化
const handleSelectinChange = (rows: any) => {
  selectedRows.value = rows;
};

// 用户点击某一行的复选框
const handleRowSelectionChange = (val: [], _row: any) => {
  // 输出当前选中的所有行数组
  console.log(val);
};

// 刷新表格
const handleRefresh = () => {
  // 清空选项
  myTableRef.value?.clearSelection();
  // 重新获取数据
  getTableData();
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
    props.delFunc!({ ids: ids.value })
      .then((res: any) => {
        if (res.data.code === 200) {
          ElMessage.success(res.data.msg);
          handleRefresh();
        } else {
          ElMessage.error(res.data.msg);
        }
      })
      .catch((err: any) => {
        ElMessage.error("数据传输失败，请检查：", err);
      });
  });
};

defineExpose({
  getTableData,
  handleDelete,
  handleRefresh,
});
</script>

<style lang="less" scoped>
.my_table {
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  width: 100%;
  height: 100%;

  &_header {
    width: 100%;
    height: 44px;
    padding: 0 6px;
    background-color: rgba(96, 122, 207, 0.9);
    color: #fff;
    display: flex;
    align-items: center;
    justify-content: space-between;

    &_title {
      font-size: 16px;
    }

    &_operation {
      flex: 2;
      display: flex;
      justify-content: flex-start;

      .operation {
        &-select {
          &-spanText {
            font-size: 12px;
            padding-left: 4px;
          }

          &-input {
            margin-left: 10px;
          }
        }
      }
    }

    &_button {
      display: flex;
      justify-content: flex-end;
    }
  }

  &_content {
    height: calc(100% - 64px - 30px - 20px);

    .table {
      width: 100%;
      height: 100%;
    }
  }

  &_page {
    height: 30px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>
