<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { selectedLocationStore, selectedTeamLocationStore } from '../../stores/selectedLocationStore';
    import { optionsStore } from "../../stores/optionsStore"
    import { teamCardsStore } from "../../stores/teamCardsStore";
    import { darkModeStore } from "../../stores/darkModeStore"
    import { storeToRefs } from 'pinia';
    import { Location } from "../../types/types"
    import { scoreBoardApi } from '../../api/api';

    //ref variables
    const { teamCards } = storeToRefs(teamCardsStore())
    const { remainingLocationOptions, allLocationOptions } = storeToRefs(optionsStore())
    const { selectedLocationName } = storeToRefs(selectedLocationStore())
    const { selectedTeam } = storeToRefs(selectedTeamLocationStore())
    const { isDarkMode } = storeToRefs(darkModeStore())

    //store methods
    const { setRemainingLocationOptions, setAllLocationOptions } = optionsStore()
    const { addTeamCard } = teamCardsStore()
    const { setSelectedTeam } = selectedTeamLocationStore()
    
    const isLoading = ref<boolean>(true)
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
        setRemainingLocationOptions(teamCards.value)
        
        //Set the current visible option for the all of the 
        // selectedLocationName.value = remainingLocationOptions.value[0]
    }

    onMounted(async () => {
        try {
            
            //If there are no locations currently already in local storage, retrieve them from the database first, and 
            //load it to reduce database load.
            if (!allLocationOptions.value.length) {
                const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")
    
                //Assign the response from the server to the above variable to be referenced later.
                locations = locationsResponse.data
                
                //Load all the locations into the following ref, storing it into local storage for faster access.
                setAllLocationOptions(locations.map(location => {          
                    return location.location_name
                }))
            }
            
            setRemainingLocationOptions(teamCards.value)

        } catch (error) {
            
        }

        isLoading.value = false
    })
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
        
        <!-- v-if="remainingLocationOptions.length"  -->
        <button 
            @click="addTeam"
            class="add-team-button"
        >Add Team</button> 
    </div>
</template>


<style scoped lang="scss">
    .team-game-location{
        border: 2px solid beige;
    }
</style>