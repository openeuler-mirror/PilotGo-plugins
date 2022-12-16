import { defineStore } from 'pinia';

export const useMacStore = defineStore('mac', {
  state: () => {
    return {
      macIp: '',
    };
  },
  getters: {},
  actions: {
    setMacIp(ip: string) {
      this.macIp = ip;
    }
  }
});
