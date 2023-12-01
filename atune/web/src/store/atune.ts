import { defineStore } from 'pinia';

export const useAtuneStore = defineStore('atune', {
  state: () => ({ count: 0 }),
  getters: {
    double: (state) => state.count * 2,
  },
  actions: {
    increment() {
      this.count++;
    },
  },
  //   persist: {
  //     enabled: true, // 开启存储
  //     strategies: [{ storage: localStorage, paths: ['atune'] }],
  //   },
});
