import { createApp } from 'vue'
import { createPinia } from 'pinia';
import { getAllLocations } from "./utils/utils"

import Toast from "vue-toastification";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import "vue-toastification/dist/index.css";

import App from './App.vue'
import router from './router/router'
import "./app.css"

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(Toast);
app.use(pinia);
app.use(router);

// await getAllLocations()

// app.mount('#app')

(async () => {
  await getAllLocations();
})();

app.mount("#app");