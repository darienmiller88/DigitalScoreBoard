<script setup lang="ts">
    import { darkModeStore } from "../../stores/darkModeStore"
    import { onMounted, ref } from 'vue';
    import { buttonActiveStore, ButtonState, GameMode } from "../../stores/buttonActiveStore"
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { selectedLocationStore, selectedTeamStore } from '../../stores/selectedLocationStore';
    import { teamCardsStore } from "../../stores/teamCardsStore";
    import { optionsStore } from "../../stores/optionsStore"; 
    import { storeToRefs } from "pinia";
    import { scoreBoardApi } from "../../api/api"
    import { Location } from "../../types/types"
    import { Icon } from "@iconify/vue"
    import Select from "../Select/Select.vue";

    //Stateful variables
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())
    const { selectedLocation, selectedLocationName } = storeToRefs(selectedLocationStore())
    const { selectedTeam, selectedTeamName, teamGameLocationName } = storeToRefs(selectedTeamStore())

    const { teamCards } = storeToRefs(teamCardsStore())
    const { allLocationOptions, remainingLocationOptions } = storeToRefs(optionsStore())

    //Stateful methods
    const { setUserCards } = scoreCardsStore()
    const { setSelectedTeam } = selectedTeamStore()
    const { setSelectedLocation } = selectedLocationStore()
    const { addTeamCard } = teamCardsStore()
    const { setRemainingLocationOptions, setAllLocationOptions } = optionsStore()

    const isLoading = ref<boolean>(true)
    let locations: Location[] = []

    const props = defineProps<{
        buttonState: GameMode
    }>()
    
    const optionClicked = async (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        try {
            const locationResponse = await scoreBoardApi.get<Location>(`/get-location/${selectedValue}`)
            
            if (currentButtonGroupState.value === ButtonState.CREATE_NEW_TEAM_GAME) {
                setSelectedTeam(locationResponse.data)
            } else if(currentButtonGroupState.value === ButtonState.CREATE_NEW_GAME){
                setUserCards(locationResponse.data.users)     
                setSelectedLocation(locationResponse.data)               
            }else{

            }

        } catch (error) {
            console.log("err in clicking option:", error);
        }
    }

    const changeTeamGameLocation = async (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        try {
            const locationResponse = await scoreBoardApi.get<Location>(`/get-location/${selectedValue}`)
            setSelectedTeam(locationResponse.data)
        } catch (error) {
            console.log("err in clicking option:", error);
        }
    }

    const getLocation = async (locationName: string): Promise<Location> => {
        try {
            const locationResponse = await scoreBoardApi.get<Location>(`/get-location/${locationName}`)
            
            return locationResponse.data
        } catch (error) {
            console.log("err in clicking option:", error);
            throw new Error("Failed to fetch location: " + error)
        }
    }

    onMounted(async () => {
        try {
            const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")

            //Assign the response from the server to the above variable to be referenced later.
            locations = locationsResponse.data

            //Take all of the names from the all of the locations, and assign them to the options 
            //variable to be listed on the dropdown menu.
            setAllLocationOptions(locations.map(location => {          
                return location.location_name
            }))      

            //IF: If there is no current location set, set it now, and the current name for the options dropdown.
            //ELSE: Assign the current name of the selectedLocation type to the selectedLocationName
            //string variable so it can show on the options dropdown.
            if (!selectedLocation.value.id) {
                selectedLocationName.value = allLocationOptions.value[0]
                setSelectedLocation(locations[0])
                console.log("selectedLocation:", selectedLocation.value);
            }else if (currentButtonGroupState.value === ButtonState.CREATE_NEW_GAME) {
                selectedLocationName.value = selectedLocation.value.location_name
                // selectedTeamGameLocationName.value = selectedLocation.value.location_name
            }else if (currentButtonGroupState.value === ButtonState.CREATE_NEW_TEAM_GAME ) {
                // selectedLocationName.value = selectedLocation.value.location_name
                teamGameLocationName.value = selectedTeam.value.location_name
            }else{
                selectedLocationName.value = remainingLocationOptions.value[0]   
            }

            //If there are no current cards set, and there is a selectedLocation, set the cards for the location.
            if (!scoreCards.value.length && selectedLocation.value) {
                setUserCards(selectedLocation.value.users)
            }

            console.log("selectedLocationName:", selectedLocationName.value);
        } catch (error) {
            console.log("err:", error)
        }

        isLoading.value = false
    })
</script>

<template>
    <div :class="`${isDarkMode ? 'dark-mode-text' : 'light-mode-text'}`">
        <div class="icon-wrapper">
            <Icon icon="svg-spinners:180-ring" v-if="isLoading"/>
        </div>

        <!-- This will show where a single player game is currently being played -->
        <div class="current-location" v-if="!isLoading && currentButtonGroupState === ButtonState.CREATE_NEW_GAME" >
            <span class="select-tag-label">Current Location:</span>

            <Select 
                :selectModel="selectedLocationName"
                :options="allLocationOptions"
                :optionClicked="optionClicked"
            />
        </div>

        <!-- If the user is creating a new team game, add this select tag. -->
        <div class="team-game-location" v-if="!isLoading && currentButtonGroupState === ButtonState.CREATE_NEW_TEAM_GAME">
            <Select 
                v-if="remainingLocationOptions.length"
                :options="remainingLocationOptions"
                :optionClicked="optionClicked"
                :selectModel="selectedTeam.location_name"
            />
            
            <!-- @click="addTeam" -->
            <button 
                v-if="remainingLocationOptions.length" 
                class="add-team-button"
            >Add Team</button> 
        </div>

        <div v-if="currentButtonGroupState === ButtonState.CREATE_NEW_TEAM_GAME">
            <span class="select-tag-label">Team game Location:</span>

            <Select 
                :selectModel="teamGameLocationName"
                :options="allLocationOptions"
                :optionClicked="optionClicked"
            />
        </div>
</div>
</template>

<style scoped lang="scss">
    // .icon-wrapper{
    //     text-align: center;
    //     font-size: 45px;
    // }

    // .select-tag-label{
        
    // }

    // .current-location{

    // }

    // .location{
    //     text-align: center;
        
    //     @media screen and (min-width: 768px) {
    //         display: flex;
    //         align-items: center;
    //         justify-content: center;
    //     }

    //     font-size: 30px;
    //     transition: 0.5s;
        
    //     .underline{
    //         text-decoration: underline;
    //     }

    //     .current-location{
    //         margin-right: 10px;
    //     }

    //     .add-team-button{
    //         padding: 10px 25px;
    //         margin: 0px 10px;

    //         border: none;
    //         border-radius: 5px;

    //         background-color: dodgerblue;
    //         color: aliceblue;
    //         font-size: 20px;
    //         transition: 0.3s;

    //         @media screen and (min-width: 768px) {
    //             &:hover{
    //                 cursor: pointer;
    //                 font-size: 26px;
    //             }
    //         }
    //     }
    // }
</style>