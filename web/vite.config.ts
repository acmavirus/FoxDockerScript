import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  base: './', // Sử dụng đường dẫn tương đối để dễ dàng nhúng
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
})
