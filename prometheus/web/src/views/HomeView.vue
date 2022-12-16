<template>
  <div style="width: 100%; height: 100%; overflow: auto; background-color: #0b0c0e;">
    <el-select v-model="macIp" @change="handleChangeIp">
      <el-option v-for="item in macIps" :key="item.labels.instance" :label="item.labels.instance"
        :value="item.labels.instance"></el-option>
    </el-select>
    <grid-layout :col-num="16" :is-draggable="grid.draggable" :is-resizable="grid.resizable" :layout="layout"
      :row-height="30" :use-css-transforms="true" :vertical-compact="true" drag-allow-from=".drag"
      drag-ignore-from=".noDrag">
      <template v-for="(item, indexVar) in layout">
        <grid-item :key="indexVar" :h="item.h" :i="item.i" :static="item.static" :w="item.w" :x="item.x" :y="item.y"
          :min-w="2" :min-h="2" @resize="SizeAutoChange(item.i, item.query.isChart)" @resized="SizeAutoChange"
          v-if="item.display">
          <div class="drag">
            <span class="drag-title">{{ item.title }}</span>
            <!-- 下拉选择 -->
            <!-- <el-dropdown class="drag-more">
              <el-icon>
                <ArrowDown />
              </el-icon>
              <template #dropdown>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item @click.native="handleEdit(item, indexVar)" :icon="Edit">Edit</el-dropdown-item>
                  <el-dropdown-item @click.native="handleDelete(indexVar)" :icon="Delete">Delete
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown> -->
          </div>
          <div class="noDrag">
            <my-echarts :ref="el => { if (el) chart[indexVar] = el }" :query="item.query"
              style="width:100%;height:100%;">
            </my-echarts>
          </div>

        </grid-item>
      </template>
    </grid-layout>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { Delete, Plus, Edit, ArrowDown } from '@element-plus/icons-vue';
import MyEcharts from '@/components/MyEcharts.vue';
import { getMacIp, getPromRules } from '@/api/prometheus';
import { useMacStore } from '@/store/mac';
import { useLayoutStore } from '@/store/charts';

const layoutStore = useLayoutStore();
let layout = reactive(layoutStore.layout_option);
const macIp = ref('');
const macIps = ref([] as any)
const chart = ref([] as any);
const grid = reactive({
  draggable: true,
  resizable: true,
  responsive: true,
});

const dialog = reactive({
  title: '',
  type: '',
  display: false
})
// 关闭dialog 
const handleCancle = (type?: string) => {
  dialog.type = '';
  dialog.title = '';
  dialog.display = false;
  if (type === 'success') {
    getMacIp();
  }
}

/* getMacIp().then(res => {
  if (res.data.url != '') {
    useMacStore().setMacIp('10.1.167.93:9100'); //res.data.reverse_dest.split('//')[1]
  }
})
getMacIp(); */
getPromRules().then(res => {
  if (res.data.status === 'success') {
    macIps.value = res.data.data && res.data.data.activeTargets;
    // 默认显示数组第一个的监控数据
    useMacStore().setMacIp(macIps.value[0].labels.instance)
  }
})
getPromRules();

const handleChangeIp = (ip: string) => {
  if (ip) {
    useMacStore().setMacIp(ip)
  }
}

// echarts大小随grid改变
const SizeAutoChange = (i: string, isChart?: boolean) => {
  if (isChart) {
    chart.value[i].resize();
  }
}

// 新增echart
const handleAdd = () => {
  dialog.type = 'add';
  dialog.title = '新增图表';
  dialog.display = true;
}

// 编辑echart
const handleEdit = (item: any, index: number) => {

}
// 删除echart
const handleDelete = (index: number) => {
  layout[index].display = false;
}

onMounted(() => {
  // 页面dom加载完成后初始化图表大小
  chart.value.forEach((item: any, index: number) => {
    chart.value[index].resize();
  });
  // layout.value = layoutStore.layout_option;
})
</script>

<style scoped lang="scss">
.vue-grid-layout {
  width: 100%;
  height: 100%;
  --title_height: 24px;


  .vue-grid-item {
    box-sizing: border-box;
    background-color: rgb(20, 22, 25);
    border: 1px solid rgb(32, 34, 38);
    border-radius: 4px;


    .drag {
      width: 100%;
      height: var(--title_height);
      border-radius: 4px 4px 0 0;
      position: absolute;
      z-index: 9999;
      display: flex;
      align-items: center;
      justify-content: center;

      &-title {
        display: flex;
        align-items: center;
        justify-content: center;
        user-select: none;
        width: 88%;
        height: 100%;
        color: rgb(187, 208, 217);
        font-size: 12px;
        font-weight: bold;

        &:hover {
          color: #fff;
          cursor: pointer;
        }
      }

      &-more {
        width: 12%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        user-select: none;
        cursor: pointer;
      }

      &:hover {
        background: rgba(32,
            34,
            38);


      }
    }

    .noDrag {
      width: 100%;
      height: calc(100% - var(--title_height));
      margin-top: var(--title_height);
      display: flex;
      justify-content: center;
      align-items: center;

      &-text {
        font-weight: bold;
        font-size: 20px;
        color: #67e0e3;
        user-select: none;
      }
    }

    &:hover {
      .vue-resizable-handle {
        background: none !important;
        width: 4px !important;
        height: 4px !important;
        bottom: 2px !important;
        right: 2px !important;
        border-right: 2px solid rgb(85, 85, 85);
        border-bottom: 2px solid rgb(85, 85, 85);
      }
    }
  }

  .vue-grid-item .resizing {
    opacity: 0.9;
  }

  .vue-grid-item .static {
    background: #cce;
  }

  .vue-grid-item .text {
    font-size: 24px;
    text-align: center;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    height: 100%;
    width: 100%;
  }

  .vue-grid-item .no-drag {
    height: 100%;
    width: 100%;
  }

  .vue-grid-item .minMax {
    font-size: 12px;
  }

  .vue-grid-item .add {
    cursor: pointer;
  }

  .vue-draggable-handle {
    position: absolute;
    width: 20px;
    height: 20px;
    top: 0;
    left: 0;
    /* background: url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='10' height='10'><circle cx='5' cy='5' r='5' fill='#999999'/></svg>") no-repeat; */
    background-color: aqua;
    background-position: bottom right;
    padding: 0 8px 8px 0;
    background-repeat: no-repeat;
    background-origin: content-box;
    box-sizing: border-box;
    cursor: pointer;
  }
}
</style>