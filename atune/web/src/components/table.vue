<template>
    <div class="my-table">
        <div class="header">
            <div class="header_content">
                <div class="table_search">
                    <slot name="table_search"></slot>
                </div>
                <div class="table_action">
                    <slot name="table_action"></slot>
                </div>
            </div>
        </div>
        <div class="content">
          <div  v-loading="loading" element-loading-text="数据加载中" element-loading-spinner="el-icon-loading">
            <el-table :data="props.tableData" @selection-change="onSelectionChange" height="100%" :header-cell-style="{ color: 'black', 'background-color': '#f6f8fd' }"
            :cell-style="{ color: 'black' }"  tooltip-effect="dark">
            <el-table-column type="selection"  width="55" align="center" ></el-table-column>
            <slot name="table"></slot>
          </el-table>
        </div>
      </div >
        <div class="pagination">
            <el-pagination 
                v-model:current-page="currentPage"  v-model:page-size="pageSize" :total="props.total" :page-sizes="props.pageSizes" 
                layout="total, sizes, prev, pager, next, jumper"
                @size-change="onPageSizeChange"
                @current-change="onCurrentPageChange">
            </el-pagination>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, toRaw } from 'vue';

const emit = defineEmits(['update:pageSize', 'update:currentPage','update:selectedData']);
let props = defineProps({
  loading:{
    type:Boolean,
    default:false
  },
  tableData: {
    type:Array,
    default:[]
  },
  searchTune: {
    type: Boolean,
    default: false
  },
  total:{
    type:Number,
    default:0
  },
  currentPage:{
    type:Number,
    default:1
  },
  pageSize:{
    type:Number,
    default:10
  },
  pageSizes: {
    type: Array,
    default: [10, 20, 50, 100],
  },
})

const currentPage = ref(props.currentPage);
const pageSize = ref(props.pageSize);

const onSelectionChange = (val: any[]) => {
    let d: any[] = []
    val.forEach((item: any) => {
        d.push(toRaw(item))
    })

    emit('update:selectedData', d)
}

const onPageSizeChange = (newSize:number) => {
    pageSize.value = newSize;
    emit('update:pageSize', newSize);
};

const onCurrentPageChange = (newPage:number) => {
    currentPage.value = newPage;
    emit('update:currentPage', newPage);
};

watch([() => props.currentPage, () => props.pageSize], ([newPage, newSize]) => {
  emit('update:currentPage', newPage);
  emit('update:pageSize', newSize);
});

// 在组件挂载时触发一次父组件传递的事件
onMounted(() => {
    emit('update:pageSize', pageSize.value);
    emit('update:currentPage', currentPage.value);
});
</script>

<style lang="less" scoped>
.my-table {
  height: 100%;
  width: 100%;
  display: flex; // 使用 flex 布局
  // overflow: visible;
  flex-direction: column;

  .header {
    width: 100%;
    height: 6%;
    border-radius: 8px 8px 0 0;
    background: linear-gradient(to right, rgb(11, 35, 117) 0%, rgb(96, 122, 207) 100%, );

    .header_content {
      height: 100%;
      margin: 0 15px;
      display: flex;
      justify-content: space-between;
      align-items: center;
      color: #fff;

      .table_search{
        margin-left: auto;
      }
    }
  }

  .content {
    // flex: 1;
    // height: 92%;
    height: calc(94% - 40px);
    overflow: scroll;
  }
  .pagination {
    width: 100%;
    height: 40px;
    margin: 0;
    :deep(.el-pagination) {
        justify-content: flex-end;
        .el-pagination__sizes {
            flex: 1,
        }
    }
    .el-pagination {
      text-align: right;
      padding-top: 5px;
      padding-bottom: 5px;

      .el-pagination__sizes,
      .el-pagination__total {
        float: left;
      }
    }
  }
}
</style>