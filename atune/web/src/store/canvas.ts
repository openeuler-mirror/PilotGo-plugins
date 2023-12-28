import { defineStore } from "pinia";
import { RectConfig } from "@/types/atune";
export const useCanvasStore = defineStore("canvas", {
  state: () => ({
    ctx: {},
    width: 0,
    fromX: 0,
    fromY: 50,
    rectWidth: 100,
    rectHeight: 50,
    rectConfig: {
      x: 0,
      y: 0,
      width: 100,
      height: 50,
      fill: "#fff",
      stroke: "#fff",
      shadowBlur: 2,
      cornerRadius: 10,
    },
  }),
  getters: {
    /**
     * 箭头函数的写法，直接把 state 中的值传进来使用
     * 如果要在 Getter 中调用其他的计算属性方法，不能使用箭头函数
     * 注意：需要自己定义当前方法的返回值类型
     * @param state
     * @returns
     */
    // 计算各个图形的起始位置
    rectStart: (state) => {
      let rectConfig: RectConfig = JSON.parse(JSON.stringify(state.rectConfig));
      rectConfig.x = state.width / 2 - state.rectWidth;
      rectConfig.y = state.fromY;
      return rectConfig;
    },
    rectPrepare: (state) => {
      let rectConfig: RectConfig = JSON.parse(JSON.stringify(state.rectConfig));
      rectConfig.x = state.width / 2 - state.rectWidth;
      rectConfig.y = state.fromY + state.fromY * 2;
      return rectConfig;
    },
    rectTune: (state) => {
      let rectConfig: RectConfig = JSON.parse(JSON.stringify(state.rectConfig));
      rectConfig.x = state.width / 2 - state.rectWidth;
      rectConfig.y = state.fromY + state.fromY * 4;
      return rectConfig;
    },
    rectRestore: (state) => {
      let rectConfig: RectConfig = JSON.parse(JSON.stringify(state.rectConfig));
      rectConfig.x = state.width / 2 - state.rectWidth;
      rectConfig.y = state.fromY + state.fromY * 6;
      return rectConfig;
    },
    arrowStartX: (state) => {
      return state.fromX - state.rectWidth / 2;
    },
    arrowStartY: (state) => {
      return state.fromY + state.rectHeight;
    },
  },
  actions: {
    setCtx(ctx: any) {
      this.ctx = ctx;
    },
    setWidth(width: number) {
      this.width = width;
      this.fromX = width / 2;
    },
    drawArrow(
      context: any,
      shape: any,
      fX: number,
      tX: number,
      fY: number,
      tY: number
    ) {
      const headlen = 10; //箭头线的长度
      const theta = 30; //箭头线与直线的夹角
      const fromX = fX;
      const fromY = fY;
      const toX = tX;
      const toY = tY;
      let arrowX, arrowY; //箭头线终点坐标
      // 计算各角度和对应的箭头终点坐标
      let angle = (Math.atan2(fromY - toY, fromX - toX) * 180) / Math.PI;
      let angle1 = ((angle + theta) * Math.PI) / 180;
      let angle2 = ((angle - theta) * Math.PI) / 180;
      let topX = headlen * Math.cos(angle1);
      let topY = headlen * Math.sin(angle1);
      let botX = headlen * Math.cos(angle2);
      let botY = headlen * Math.sin(angle2);
      context.beginPath();
      // draw a line
      context.moveTo(fromX, fromY);
      context.lineTo(toX, toY);
      // draw top arrow
      arrowX = toX + topX;
      arrowY = toY + topY;
      context.moveTo(arrowX, arrowY);
      context.lineTo(toX, toY);
      // draw bottom arrow
      arrowX = toX + botX;
      arrowY = toY + botY;
      context.lineTo(arrowX, arrowY);
      context.closePath();

      // special Konva.js method
      context.fillStrokeShape(shape);
    },
  },
});
