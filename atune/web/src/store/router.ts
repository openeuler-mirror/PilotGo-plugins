import { defineStore } from "pinia";
import { Task, Atune } from "@/types/atune";

export const useRouterStore = defineStore("router", {
  state: () => ({
    route: "",
  }),
  getters: {
    currentRoute: (state) => state.route,
  },
  actions: {
    setCurrentRoute(route: string) {
      this.route = route;
    },
    showRoute(path: string, str: string) {
      return path.indexOf(str) >= 0;
    },
  },
});
