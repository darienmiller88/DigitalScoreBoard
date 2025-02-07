import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Location } from "../types/types"

export const selectedLocationStore = defineStore("selectedLocation", () => {
    const selectedLocation = ref<Location>()

    const setSelectedLocation = (location: Location) => {
        selectedLocation.value = location
    }

    return { selectedLocation, setSelectedLocation }
}, {
    persist: true
})