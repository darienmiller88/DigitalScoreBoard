import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Team } from "../types/types"

export const optionsStore = defineStore("optionsStore", () => {
    //When a tournament is being held, client can choose this option.
    const selectAllSites = "Select from all sites"

    //This one only contains the 9 ADAPT sites to choose from
    const allLocationOptions = ref<string[]>([])

    //This one contains the 9 ADAPT sites, plus the constant above so the client can choose players
    //from all sites, not just one.
    const allLocationOptionsForTournament =  ref<string[]>([])

    //This will hold the players who have not yet been chosen for a game.
    const remainingLocationOptions = ref<string[]>([])

    //Overides the Location allLocationOptions with a set of new ones.
    const setAllLocationOptions = (newOptions: string[]) => {
        allLocationOptions.value = newOptions
        allLocationOptionsForTournament.value = [selectAllSites, ...newOptions]
    }

    //Removes a number of teams (Locations really) from the list of all locations, returning the remaining teams.
    const setRemainingLocationOptions = (teams: Team[]) => {
        remainingLocationOptions.value = allLocationOptions.value.filter(
            location => !teams.some(team => team.team_name === location)
        )
    }

    //add a team back to the list of Remaining teams. This happens when the "clear button" is clicked on a Team card.
    const addOptionToRemainingLocationOptions = (option: string) => {
        remainingLocationOptions.value.push(option)
    }
   
    return { 
        allLocationOptions, 
        allLocationOptionsForTournament,
        selectAllSites,
        remainingLocationOptions, 
        setAllLocationOptions, 
        setRemainingLocationOptions,
        addOptionToRemainingLocationOptions
    }
}, {
    persist: true
})