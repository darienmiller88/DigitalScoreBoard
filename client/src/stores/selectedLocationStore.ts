import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Location } from "../types/types"

export const selectedLocationStore = defineStore("selectedLocation", () => {
    const selectedLocationName = ref<string>("")
    const selectedTeamGameLocationName = ref<string>("")
    const selectedLocation = ref<Location>({
        id: "",
        users: [],
        location_name: ""
    })

    const setSelectedLocation = (location: Location) => {
        selectedLocation.value = location
        selectedLocationName.value = location.location_name
    }

    return { selectedLocation, selectedLocationName, selectedTeamGameLocationName, setSelectedLocation }
}, {
    persist: true
})