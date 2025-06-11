import { defineStore } from 'pinia';
import { ref } from 'vue';

export const darkModeStore = defineStore("dark-mode", () => {
    const isDarkMode = ref<boolean>(true)

    const toggleDarkMode = () => {
        isDarkMode.value = !isDarkMode.value
    }
    
    return { isDarkMode, toggleDarkMode }
}, {
    persist: true
})