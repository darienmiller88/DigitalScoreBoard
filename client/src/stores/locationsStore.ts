import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Location } from "../types/types"

export const locationsStore = defineStore("locationsStore", () => {
    const locations = ref<Location[]>([])

    const getLocations = (): Location[] =>{
        return locations.value
    }
   
    return { getLocations }
}, {
    persist: true
})