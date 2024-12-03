/* 
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: 赵振芳 <zhaozhenfang@kylinos.cn>
 * Date: Mon Nov 25 17:07:09 2024 +0800
 */
import { createApp } from 'vue'
import './style.css'

import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue'

export const app = createApp(App)
app.use(ElementPlus)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.mount('#app')