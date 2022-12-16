import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import VueGridLayout from 'vue-grid-layout';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import './styles/main.scss'
import '@/utils/echarts'

const app = createApp(App);

const pinia = createPinia();

app.use(pinia);
app.use(router);
app.use(VueGridLayout)
app.use(ElementPlus)
// 引入element-plus图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
app.mount('#app')