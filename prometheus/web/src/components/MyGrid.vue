<template>
  <div style="width: 100%; height: 100%; overflow: auto;">
    <grid-layout :col-num="12" :is-draggable="grid.draggable" :is-resizable="grid.resizable" :layout="layout"
      :row-height="30" :use-css-transforms="true" :vertical-compact="true" drag-allow-from=".drag"
      drag-ignore-from=".noDrag">
      <grid-item v-for="(item, indexVar) in layout" :key="indexVar" :h="item.h" :i="item.i" :static="item.static"
        :w="item.w" :x="item.x" :y="item.y" @resize="SizeAutoChange" @resized="SizeAutoChange">
        <div class="drag">
          <span class="title">{{item.title}}</span>
        </div>
        <div class="noDrag">
          <my-echarts v-if="item.isChart" :ref="el => { if (el) chart[indexVar] = el }" :options="item.option"
            style="padding-top: 30px;">
          </my-echarts>
          <span v-else class="noDrag-text">{{item.value}}</span>
        </div>

      </grid-item>
    </grid-layout>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import MyEcharts from './MyEcharts.vue';
const options = reactive({
  xAxis: {
    type: 'category',
    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      data: [120, 200, 150, 80, 70, 110, 130],
      type: 'bar',
      showBackground: true,
      backgroundStyle: {
        color: 'rgba(180, 180, 180, 0.2)'
      }
    }
  ]
})

onMounted(() => {
  chart.value.forEach((item: any, index: number) => {
    chart.value[index].resize();
  });
})
const chart = ref([] as any);
const SizeAutoChange = (i: string, x: number, y: number, newH: number, newW: number) => {
  chart.value[i].resize();
}


const grid = reactive({
  draggable: true,
  resizable: true,
  responsive: true,
});
const layout = [
  { x: 0, y: 0, w: 2, h: 2, i: '0', static: true, title: 'static', isChart: false, value: '1.2day' },
  { x: 2, y: 0, w: 2, h: 4, i: '1', display: true, title: 'cpu', isChart: false, value: 4 },
  { x: 4, y: 0, w: 2, h: 5, i: '2', isChart: false, value: '2.35GiB' },
  { x: 6, y: 0, w: 2, h: 3, i: '3', option: network_opt },
  { x: 8, y: 0, w: 2, h: 3, i: '4' },
  { x: 10, y: 0, w: 2, h: 3, i: '5' },
  { x: 0, y: 5, w: 2, h: 5, i: '6' },
  { x: 2, y: 5, w: 2, h: 5, i: '7' },
  { x: 4, y: 5, w: 2, h: 5, i: '8' },
  { x: 6, y: 4, w: 2, h: 4, i: '9' },
  { x: 8, y: 4, w: 2, h: 4, i: '10' },
  { x: 10, y: 4, w: 2, h: 4, i: '11' },
  { x: 0, y: 10, w: 2, h: 5, i: '12' },
  { x: 2, y: 10, w: 2, h: 5, i: '13' },
  { x: 4, y: 8, w: 2, h: 4, i: '14' },
  { x: 6, y: 8, w: 2, h: 4, i: '15' },
  { x: 8, y: 10, w: 2, h: 5, i: '16' },
  { x: 10, y: 4, w: 2, h: 2, i: '17' },
  { x: 0, y: 9, w: 2, h: 3, i: '18' },
  { x: 2, y: 6, w: 2, h: 2, i: '19' },
]
const layoutHeight = 130;
const layoutConfig = {
  height: 330, // 默认高度
  dialogVisible: false // 是否可拖拽或改变大小
}

</script>

<style scoped lang="scss">
.vue-grid-layout {
  width: 100%;
  height: 100%;
}

.vue-grid-item {
  box-sizing: border-box;
  background-color: rgb(20, 22, 25);
  border: 1px solid rgb(32, 34, 38);
  border-radius: 4px;


  .drag {
    width: 100%;
    height: 30px;
    border-radius: 4px 4px 0 0;
    position: absolute;
    z-index: 9999;

    .title {
      display: flex;
      align-items: center;
      justify-content: center;
      user-select: none;
      width: 30%;
      margin: 0 auto;
      height: 100%;
      color: rgb(187, 208, 217);
      font-size: 16px;
      font-weight: bold;

      &:hover {
        color: #fff;
        cursor: pointer;
      }
    }


    &:hover {
      background: rgba(32,
          34,
          38);
    }
  }

  .noDrag {
    width: 100%;
    height: 100%;

    &-text {
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 16px;
      color: rgb(48, 162, 38);
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
</style>