import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss()],
  build: {
    rollupOptions: {
      external: ['pdfjs-dist/legacy/build/pdf', 'pdfjs-dist/types/src/display/api'],
    },
  },
});
