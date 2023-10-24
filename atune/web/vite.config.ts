import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path';


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src') //在任何模块文件内部，可以使用__dirname变量获取当前模块文件所在目录的完整绝对路径。
    }
  },
  server:{
    host:'0.0.0.0',
    port:8080
  },
})

