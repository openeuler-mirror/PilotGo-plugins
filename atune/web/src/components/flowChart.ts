import { useCanvasStore } from "@/store/canvas";
const fromX = useCanvasStore().width || 0;
console.log("起始位置X:", fromX);
export const rectStart = {
  x: fromX / 2,
  y: 50,
  width: 100,
  height: 40,
  fill: "#fff",
  stroke: "#fff",
  shadowBlur: 2,
  cornerRadius: 10,
};

export const rectPrepare = {
  x: 20,
  y: 50,
  width: 100,
  height: 40,
  fill: "#fff",
  stroke: "#fff",
  shadowBlur: 2,
  cornerRadius: 10,
};

export const rectTune = {
  x: 20,
  y: 50,
  width: 100,
  height: 40,
  fill: "#fff",
  stroke: "#fff",
  shadowBlur: 2,
  cornerRadius: 10,
};

export const rectRestore = {
  x: 20,
  y: 50,
  width: 100,
  height: 40,
  fill: "#fff",
  stroke: "#fff",
  shadowBlur: 2,
  cornerRadius: 10,
};
