import { createApp } from 'vue'
import App from './App.vue'
import "./assets/css/main.css";

import '@mdi/font/css/materialdesignicons.min.css'
import './assets/tailwind.css'
import router from './router/router'

createApp(App).use(router).mount('#app')
