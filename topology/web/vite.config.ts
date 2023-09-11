import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server:{
    proxy:{
      "/plugin/api": {
        target: 'http://192.168.241.129:9991',
        changeOrigin:true,
        // rewrite: (path)=> path.replace("/^\/api/", ""),
      }
    }
  }
})
