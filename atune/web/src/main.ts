import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import ElementPlus from "element-plus";
import MyTable from "@/components/table.vue";
import MyButton from "@/components/myButton.vue";
import "element-plus/dist/index.css";
import router from "./router";
import { createPinia } from "pinia";
import VueKonva from "vue-konva";

const app = createApp(App);
// 设置全局变量
// app.config.globalProperties.$router = router;
app.component("my-table", MyTable);
app.component("my-button", MyButton);

app.use(createPinia());
app.use(router);
app.use(ElementPlus);
app.use(VueKonva);

app.mount("#app");
