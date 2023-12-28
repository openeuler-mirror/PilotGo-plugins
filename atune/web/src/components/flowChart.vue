<template>
  <div id="flowChart">
    <v-stage ref="stage" :config="stageSize" v-if="showFlow">
      <v-layer>
        <v-rect
          @click="handleClick('reStart')"
          :config="reStart"
          @mousemove="handleMouseMove('reStart')"
          @mouseout="handleMouseOut('reStart')"
        /><v-rect
          :config="rectPrepare"
          @click="handleClick('rectPrepare')"
          @mousemove="handleMouseMove('rectPrepare')"
          @mouseout="handleMouseOut('rectPrepare')"
        /><v-rect
          :config="rectTune"
          @click="handleClick('rectTune')"
          @mousemove="handleMouseMove('rectTune')"
          @mouseout="handleMouseOut('rectTune')"
        /><v-rect
          :config="rectRestore"
          @click="handleClick('rectRestore')"
          @mousemove="handleMouseMove('rectRestore')"
          @mouseout="handleMouseOut('rectRestore')"
        />
        <v-text :config="{ x: 0, y: 20, text: 'start', fontSize: 16 }" />
        <v-text :config="{ x: 20, y: 40, text: 'prepare', fontSize: 16 }" />
        <v-text :config="{ x: 20, y: 80, text: 'restore', fontSize: 16 }" />
        <v-text :config="{ x: 20, y: 60, text: 'tune', fontSize: 16 }" />
        <v-shape
          :config="{
            sceneFunc: arrowConfig1,
            fill: '#00D2FF',
            stroke: '#00D2FF',
            strokeWidth: 2,
          }"
        />
        <v-shape
          :config="{
            sceneFunc: arrowConfig2,
            fill: '#00D2FF',
            stroke: '#00D2FF',
            strokeWidth: 2,
          }"
        />
      </v-layer>
      <v-layer ref="dragLayer"></v-layer>
    </v-stage>
  </div>
</template>
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { useCanvasStore } from "@/store/canvas";
import { RectConfig } from "@/types/atune";
import { pa } from "element-plus/es/locale/index.mjs";
const showFlow = ref(false);
const stage = ref();
const stageW = ref(0);
const stageH = ref(0);
const stageSize = ref({});
const reStart = ref({} as RectConfig);
const rectPrepare = ref({} as RectConfig);
const rectTune = ref({} as RectConfig);
const rectRestore = ref({} as RectConfig);

let arrowStartX = 0;
let arrowStartY = 0;
let rectHeight = 0;
onMounted(() => {
  const stageDiv = document.getElementById("flowChart");
  stageW.value = stageDiv!.clientWidth;
  stageH.value = stageDiv!.clientHeight;
  useCanvasStore().setWidth(stageW.value);
});
// 舞台的大小
setTimeout(() => {
  const width = stageW.value;
  const height = stageH.value;
  stageSize.value = {
    width: width,
    height: height,
  };
  reStart.value = useCanvasStore().rectStart;
  rectPrepare.value = useCanvasStore().rectPrepare;
  rectTune.value = useCanvasStore().rectTune;
  rectRestore.value = useCanvasStore().rectRestore;

  arrowStartX = useCanvasStore().arrowStartX;
  arrowStartY = useCanvasStore().arrowStartY;
  rectHeight = useCanvasStore().rectHeight;
  showFlow.value = true;
}, 100);

// 鼠标移入事件
const handleMouseMove = (shapeName: string) => {
  console.log("移入", shapeName);
  eval(shapeName).value.fill = "yellow";
};
// 鼠标移出事件
const handleMouseOut = (shapeName: string) => {
  console.log("移出", shapeName);
  eval(shapeName).value.fill = "#fff";
};
// 鼠标点击事件
const handleClick = (shapeName: string) => {
  console.log("点击了：", shapeName);
};

// 配置箭头参数
let arrowPosition = (idNum: number) => {
  let arrowFX = arrowStartX;
  let arrowTX = arrowStartX;
  let arrowFY = arrowStartY + rectHeight * idNum + 1;
  let arrowTY = arrowStartY + rectHeight * (idNum + 1) - 2;
  return { arrowFX, arrowTX, arrowFY, arrowTY };
};
const arrowConfig1 = (context: any, shape: any) => {
  let { arrowFX, arrowTX, arrowFY, arrowTY } = arrowPosition(0);

  useCanvasStore().drawArrow(
    context,
    shape,
    arrowFX,
    arrowTX,
    arrowFY,
    arrowTY
  );
};
const arrowConfig2 = (context: any, shape: any) => {
  useCanvasStore().drawArrow(
    context,
    shape,
    arrowStartX,
    arrowStartX,
    arrowStartY + rectHeight * 2 + 1,
    arrowStartY + rectHeight * 3 - 2
  );
};
</script>
<style lang="less" scoped>
#flowChart {
  width: 100%;
  height: 100%;
}
</style>
