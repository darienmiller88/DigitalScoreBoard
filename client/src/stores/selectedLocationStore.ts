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

//For team player tournament games, this will be the store to hold the current game state.
export const selectedTeamLocationStore = defineStore("selectedTeamLocation", () => {
    // This stores the name of the team that was selected. This will be used to bind to the options tag.
    const selectedTeamName = ref<string>("")

    // This stores the ADAPT location of where the team game is being played.
    const teamGameLocationName = ref<string>("")

    //Finally, this stores the location object of the team game. It will be referenced when the user wants to add a
    //team to the current list of teams, so instead of making a database call, it will done here.
    const selectedTeam = ref<Location>({
        id: "",
        users: [],
        location_name: ""
    })

    const setTeamGameLocation = (locationName: string) => {
        teamGameLocationName.value = locationName
    }

    const setSelectedTeam = (location: Location) => {
        selectedTeam.value = location
    }

    const setSelectedTeamName = (teamName: string) => {
        selectedTeamName.value = teamName
    }

    return { setTeamGameLocation, setSelectedTeam, setSelectedTeamName, teamGameLocationName, selectedTeamName, selectedTeam }
}, {
    persist: true
})