import { defineStore } from "pinia";
import { Task, Atune } from "@/types/atune";

export const useAtuneStore = defineStore("atune", {
  state: () => ({
    count: 0,
    taskRow: {} as Task,
    tuneRow: {} as Atune,
  }),
  getters: {
    double: (state) => state.count * 2,
  },
  actions: {
    increment() {
      this.count++;
    },
    // 设置taskRow数据
    setTaskRow(row: Task) {
      this.taskRow = row;
    },
    // 设置taskRow数据
    setTuneRow(row: Atune) {
      this.tuneRow = row;
      console.log(this.tuneRow);
    },
  },
  //   persist: {
  //     enabled: true, // 开启存储
  //     strategies: [{ storage: localStorage, paths: ['atune'] }],
  //   },
});
