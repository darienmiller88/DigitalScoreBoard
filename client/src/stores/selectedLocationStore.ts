import { defineStore } from 'pinia';
import { ref } from 'vue';

export const selectedLocationStore = defineStore("selectedLocation", () => {
    const selectedLocation = ref<string>("")

    const setSelectedLocation = (location: string) => {
        selectedLocation.value = location
    }

    return { selectedLocation, setSelectedLocation }
}, {
    persist: true
})