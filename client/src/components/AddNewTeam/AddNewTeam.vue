<script setup lang="ts">
    import { ref } from 'vue';
    import { selectedLocationStore, selectedTeamLocationStore } from '../../stores/selectedLocationStore';
    import { optionsStore } from "../../stores/optionsStore"
    import { teamCardsStore } from "../../stores/teamCardsStore";
    import { darkModeStore } from "../../stores/darkModeStore"
    import { storeToRefs } from 'pinia';
    import { Location } from "../../types/types"
    import { scoreBoardApi } from '../../api/api';

    //ref variables
    const { teamCards } = storeToRefs(teamCardsStore())
    const { remainingLocationOptions } = storeToRefs(optionsStore())
    const { selectedLocationName } = storeToRefs(selectedLocationStore())
    const { selectedTeam } = storeToRefs(selectedTeamLocationStore())
    const { isDarkMode } = storeToRefs(darkModeStore())

    //store methods
    const { setRemainingLocationOptions} = optionsStore()
    const { addTeamCard } = teamCardsStore()
    const { setSelectedTeam } = selectedTeamLocationStore()
    
    const isLoading = ref<boolean>(false)
    let locations: Location[] = []

    const optionClicked = async (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        try {
            const locationResponse = await scoreBoardApi.get<Location>(`/get-location/${selectedValue}`)
            
            setSelectedTeam(locationResponse.data)
        } catch (error) {
            console.log("err in clicking option:", error);
        }
    }

    const addTeam = () => {

        //Create a new anonymous Team card with the following data:
        //The location for which team is playing, a score of 0, and no players (yet).
        addTeamCard({
            team_name: selectedLocationName.value,
            score: 0,
            players: []
        })

        //After adding the card, remove it from options.
        setRemainingLocationOptions(locations.filter(
            location => !teamCards.value.some(team => team.team_name === location.location_name)
        ).map(location => location.location_name))
        
        //Set the current visible option for the all of the 
        // selectedLocationName.value = remainingLocationOptions.value[0]
    }
</script>

<template>
    <div :class="`${isDarkMode ? 'dark-mode-text' : 'light-mode-text'}`">
    </div>
    <div class="team-game-location" v-if="!isLoading">
        <Select 
            v-if="remainingLocationOptions.length"
            :options="remainingLocationOptions"
            :optionClicked="optionClicked"
            :selectModel="selectedTeam"
        />
        
        <button 
            @click="addTeam"
            v-if="remainingLocationOptions.length" 
            class="add-team-button"
        >Add Team</button> 
    </div>
</template>


<style scoped lang="scss">
    .team-game-location{
        border: 2px solid beige;
    }
</style>