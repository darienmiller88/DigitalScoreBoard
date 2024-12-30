import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Location } from "../types/types"

export const locationsStore = defineStore("locationsStore", () => {
    const locationsFromLocalStorage = ref<Location[]>([])

    const getLocations = (): Location[] =>{
        return locationsFromLocalStorage.value
    }

    const addLocation = (location: Location) => {
        locationsFromLocalStorage.value.push(location)
    }
   
    return { locationsFromLocalStorage, getLocations, addLocation }
}, {
    persist: true
})