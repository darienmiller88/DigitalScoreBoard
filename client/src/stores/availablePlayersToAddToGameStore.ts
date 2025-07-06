import { defineStore } from 'pinia';
import { ref } from 'vue';

export const availablePlayersToAddStore = defineStore("availablePlayersToAddStore", () => {
    const availablePlayersToAdd = ref<string[]>([])

    
    return { availablePlayersToAdd }
}, {
    persist: true
})