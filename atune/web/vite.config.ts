import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  base: '/plugin/atune/',
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.join(__dirname, 'src'), //在任何模块文件内部，可以使用__dirname变量获取当前模块文件所在目录的完整绝对路径。
    },
  },
  server: {
    port: 8080,
    https: false,
    cors: true,
    proxy: {
      '/': {
        target: 'http://localhost:8099',
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: 'dist',
    assetsDir: 'static/', //指定静态资源存放路径
    sourcemap: false, //是否构建source map 文件
    chunkSizeWarningLimit: 1500, //chunk 大小警告的限制，默认500KB
  },
});
