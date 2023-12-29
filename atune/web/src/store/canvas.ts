import { defineStore } from "pinia";
import { RectConfig } from "@/types/atune";
export const useCanvasStore = defineStore("canvas", {
  state: () => ({
    ctx: {},
    width: 0,
    fromX: 0,
    fromY: 50,
    arrowLength: 50,
    rectWidth: 100,
    rectHeight: 50,
    rectConfig: {
      x: 0,
      y: 0,
      width: 100,
      height: 50,
      stroke: "#fff",
      shadowBlur: 2,
      cornerRadius: 10,
    },
    textConfig: { x: 0, y: 20, text: "start", fontSize: 16 },
  }),
  getters: {
    /**
     * 箭头函数的写法，直接把 state 中的值传进来使用
     * 如果要在 Getter 中调用其他的计算属性方法，不能使用箭头函数
     * 注意：需要自己定义当前方法的返回值类型
     * @param state
     * @returns
     */
    // 箭头的起始位置X Y
    arrowStartX: (state) => {
      return state.fromX;
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
    drawRect(idNum: number) {
      // 画矩形
      let rectConfig: RectConfig = JSON.parse(JSON.stringify(this.rectConfig));
      rectConfig.x = this.width / 2 - this.rectWidth / 2;
      rectConfig.y = this.fromY + this.fromY * idNum * 2;
      return rectConfig;
    },
    // 计算文字像素
    getTextWidth(text: string, fontSize: number, fontWeight: string) {
      // 创建临时元素
      const ele: HTMLElement = document.createElement("div");
      ele.style.position = "absolute";
      ele.style.whiteSpace = "nowrap";
      ele.style.fontSize = fontSize + "px";
      ele.style.fontWeight = fontWeight;
      ele.innerText = text;
      document.body.append(ele);
      const width: number = ele.getBoundingClientRect().width;
      document.body.removeChild(ele);
      return width;
    },
    // 写text
    writeText(idNum: number, text: string) {
      let getTextWidth = this.getTextWidth(text, 16, "none");
      let textConfig = JSON.parse(JSON.stringify(this.textConfig));
      textConfig.x = this.fromX - getTextWidth / 2;
      textConfig.y =
        this.fromY + this.fromY * idNum * 2 + this.rectHeight / 2 - 6;
      textConfig.text = text;
      return textConfig;
    },
    drawArrow(
      // 画箭头
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
