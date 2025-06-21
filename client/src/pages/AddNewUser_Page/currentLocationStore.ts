import { defineStore } from 'pinia';
import { ref } from 'vue';

export const currentLocationStore = defineStore("currentLocation", () => {
    const currentLocation = ref<string>("")

    const setCurrentLocation = (location: string) => {
        currentLocation.value = location
    }

    return { setCurrentLocation, currentLocation }
}, {
    persist: true
})