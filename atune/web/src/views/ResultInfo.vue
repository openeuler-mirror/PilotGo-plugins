<template>
  <div class="container">
    <div class="top">
        <span class="top-title">A-Tune调优结果展示</span>
    </div>
  <div
      class="table" >
    <my-table 
        :tableData="tableData"
        :loading="loading"
        :total="total" 
        :pageSize.sync="pageSize" 
        :currentPage.sync="currentPage"
        @update:pageSize="onPageSizeChange" 
        @update:currentPage="onCurrentPageChange"
        @update:selectedData="onSelectionChange">
      <template v-slot:table_search>
        <el-input placeholder="请输入IP地址进行搜索..." :prefix-icon="Search" clearable
          style="width: 280px;margin-right: 10px;" v-model="searchIP" @keydown.enter.native="searchData"></el-input>
        <el-button :icon="Search" @click="searchData">搜索</el-button>
        <el-button class="delete-button" @click="handleDelete">删除</el-button>
      </template>
      <template v-slot:table :loading="loading">
        <el-table-column prop="machine_ip" label="IP"> </el-table-column>
        <el-table-column prop="command" label="执行命令"></el-table-column>
        <el-table-column prop="stdout" label="结果"></el-table-column>
        <el-table-column prop="resError" label="接口错误"></el-table-column>
      </template>
    </my-table> 
  </div> 
  </div>
    
</template>

<script lang="ts" setup>
import {  onMounted, ref} from 'vue';
import  MyTable  from "@/components/table.vue";
import { Search } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus';
import { getResultLists, searchResult, deleteResult} from '@/api/result'

const searchIP = ref("")
const tableData = ref([])
const total=ref(0)
const currentPage = ref(1);
const pageSize = ref(10);
const loading=ref(false)
const selectedRows = ref<any[]>([])

// 获取执行结果
function getData() {
  getResultLists({
      page: currentPage.value,
      size: pageSize.value,
    }).then(res=>{
      if (res.data.code===200){
        currentPage.value=res.data.page
        pageSize.value=res.data.size
        tableData.value = res.data.data;
        total.value = res.data.total;
        loading.value = false;
      } else {
            ElMessage.error("failed to get results info: " + res.data.msg)
        }
    }).catch((err: any) => {
        ElMessage.error("failed to get results info:" + err.msg)
    })
}

// 高级搜索执行结果
function searchData(){
  searchResult({
    page: currentPage.value,
    size: pageSize.value,
    searchKey: searchIP.value
  }).then(res => {
    if (res.data.code===200){
      currentPage.value=res.data.page
      pageSize.value=res.data.size
      tableData.value = res.data.data;
      total.value = res.data.total;
      loading.value = false;
    }else {
            ElMessage.error("failed to search results info: " + res.data.msg)
        }
  }).catch(error => {
    ElMessage.error('Error search data:', error);
  })
}

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
      deleteResult({ ids: ids.value }).then(res => {
        if (res.data.code === 200) {
          getData()
          ElMessage.success(res.data.msg)
        } else {
          ElMessage.error(res.data.msg)
        }
      }).catch(err => {
        ElMessage.error("数据传输失败，请检查：", err)
      });
    })
}

// 选中多选框
const onSelectionChange = (selection: any[]) => {
  selectedRows.value = selection;
};

const onPageSizeChange = (newSize:number) => {
  pageSize.value = newSize;
  getData();
};

const onCurrentPageChange = (newPage:number) => {
  currentPage.value = newPage;
  getData();
};

onMounted(()=>{
  loading.value = true;
  getData();
})
</script>
<style lang="less" scoped>
.container{
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
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
  .table{
    flex: 1;
    display: flex;
    width: 100%;
    // height: ;
  }

}
</style>