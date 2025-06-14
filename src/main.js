import { createApp } from 'vue'
import App from './App.vue'
import {i18n} from "@/i18n/i18n.js";

createApp(App).use(i18n).mount('#app')
