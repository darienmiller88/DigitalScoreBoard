import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Team } from "../types/types"

export const optionsStore = defineStore("optionsStore", () => {
    const allLocationOptions = ref<string[]>([])
    const remainingLocationOptions = ref<string[]>([])

    //Overides the Location allLocationOptions with a set of new ones.
    const setAllLocationOptions = (newOptions: string[]) => {
        allLocationOptions.value = newOptions
    }

    //Removes a number of teams (Locations really) from the list of all locations, returning the remaining teams.
    const setRemainingLocationOptions = (teams: Team[]) => {
        remainingLocationOptions.value = allLocationOptions.value.filter(
            location => !teams.some(team => team.team_name === location)
        )
    }

    //Overides the remainingLocationOptions options with a set of locations not added to a team game.
    // const setRemainingLocationOptions = (newOptions: string[]) => {
    //     remainingLocationOptions.value = newOptions
    // }

    //add a team back to the list of Remaining teams. This happens when the "clear button" is clicked on a Team card.
    const addOptionToRemainingLocationOptions = (option: string) => {
        remainingLocationOptions.value.push(option)
    }
   
    return { 
        allLocationOptions, 
        remainingLocationOptions, 
        setAllLocationOptions, 
        // setRemainingLocationOptions,
        setRemainingLocationOptions,
        addOptionToRemainingLocationOptions
    }
})