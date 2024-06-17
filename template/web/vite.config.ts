import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  base: "/plugin/template",
  plugins: [vue()],
  server: {
    proxy: {
      '/plugin/template/api': {
        target: 'https://10.41.107.29:49151',
        secure:false,
        changeOrigin: true,
        rewrite: path => path.replace(/^\//, '')
      },
    },
  }
})
