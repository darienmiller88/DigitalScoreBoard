import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Location } from "../types/types"

export const locationsStore = defineStore("locationsStore", () => {
    const locationsFromLocalStorage = ref<Location[]>([])

    const addLocation = (location: Location) => {
        locationsFromLocalStorage.value.push(location)
    }
   
    return { locationsFromLocalStorage, addLocation }
}, {
    persist: true
})