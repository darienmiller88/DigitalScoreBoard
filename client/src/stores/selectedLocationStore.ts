import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Location } from "../types/types"

//For single player games, this stores the location of where the game is being played.
export const selectedLocationStore = defineStore("selectedLocation", () => {
    const selectedLocationName = ref<string>("")
    const selectedLocation = ref<Location>({
        id: "",
        users: [],
        location_name: ""
    })

    const setSelectedLocation = (location: Location) => {
        selectedLocation.value = location
        selectedLocationName.value = location.location_name
    }

    return { selectedLocation, selectedLocationName, setSelectedLocation }
}, {
    persist: true
})

export const selectedTeamLocationStore = defineStore("selectedTeamLocation", () => {
    // This stores the name of the team that was selected.
     const selectedTeamGameLocationName = ref<string>("")

    // This stores the location of where the team game is being played.
    const teamGameLocationName = ref<string>("")

    //Finally, this stores the location object of the team game.
    const selectedTeamGameLocation = ref<Location>({
        id: "",
        users: [],
        location_name: ""
    })

    const setSelectedTeamLocation = (location: Location) => {
        selectedTeamGameLocation.value = location
    }

    return { selectedTeamGameLocation, teamGameLocationName, selectedTeamGameLocationName, setSelectedTeamLocation }
}, {
    persist: true
})