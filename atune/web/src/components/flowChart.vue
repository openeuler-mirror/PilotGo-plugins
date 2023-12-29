<template>
  <div id="flowChart">
    <v-stage ref="stage" :config="stageSize" v-if="showFlow">
      <v-layer>
        <v-rect
          v-for="(item, index) in rectArr"
          @click="handleClick(index, item)"
          :config="{
            ...useCanvasStore().drawRect(index),
            fill: rectBgColor[index],
          }"
        />
        <v-text
          v-for="(item, index) in rectArr"
          @click="handleClick(index, item)"
          :config="useCanvasStore().writeText(index, item)"
        />

        <v-shape
          v-for="item in 3"
          :config="{
            sceneFunc: (...event:any) =>{arrowConfig(event[0],event[1], item)},
            fill: '#222',
            stroke: '#222',
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
const showFlow = ref(false);
const stage = ref();
const stageW = ref(0);
const stageH = ref(0);
const stageSize = ref({});
const rectBgColor = ref(["yellow", "#fff", "#fff", "#fff"]);
const emit = defineEmits(["clickRect"]);
let rectArr = ["start", "prepare", "restore", "tune"];
let arrowStartX = 0;
let arrowStartY = 0;
let arrowLength = 0;
onMounted(() => {
  const stageDiv = document.getElementById("flowChart");
  stageW.value = stageDiv!.clientWidth;
  stageH.value = stageDiv!.clientHeight;
  useCanvasStore().setWidth(stageW.value);
});
// 舞台的大小，基本参数的动态获取
setTimeout(() => {
  const width = stageW.value;
  const height = stageH.value;
  stageSize.value = {
    width: width,
    height: height,
  };
  arrowStartX = useCanvasStore().arrowStartX;
  arrowStartY = useCanvasStore().arrowStartY;
  arrowLength = useCanvasStore().arrowLength;
  showFlow.value = true;
}, 100);

// 鼠标移入事件
const handleMouseMove = (_shapeName: string) => {};
// 鼠标移出事件
const handleMouseOut = (_shapeName: string) => {};
// 鼠标点击事件
const handleClick = (index: number, rectText: string) => {
  rectBgColor.value[index] = "yellow";
  rectBgColor.value.forEach((_item, itemIndex) => {
    if (itemIndex != index) {
      rectBgColor.value[itemIndex] = "#fff";
    }
  });
  emit("clickRect", rectText);
};

// 配置箭头参数
let arrowPosition = (idNum: number) => {
  let arrowFX = arrowStartX;
  let arrowTX = arrowStartX;
  let arrowFY = arrowStartY * idNum + 1;
  let arrowTY = arrowStartY * idNum + arrowLength - 2;
  return [arrowFX, arrowTX, arrowFY, arrowTY];
};
const arrowConfig = (context: any, shape: any, idNum: number) => {
  let params: number[] = arrowPosition(idNum);
  useCanvasStore().drawArrow(
    context,
    shape,
    params[0],
    params[1],
    params[2],
    params[3]
  );
};
</script>
<style lang="less" scoped>
#flowChart {
  width: 100%;
  height: 100%;
}
</style>
