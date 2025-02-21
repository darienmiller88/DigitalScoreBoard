<script setup lang="ts">
    import { darkModeStore } from "../../stores/darkModeStore"
    import { onMounted, ref } from 'vue';
    import { buttonActiveStore, ButtonState } from "../../stores/buttonActiveStore"
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
    import { teamCardsStore } from "../../stores/teamCardsStore";
    import { storeToRefs } from "pinia";
    import { scoreBoardApi } from "../../api/api"
    import { Location } from "../../types/types"
    import { Icon } from "@iconify/vue"

    //Stateful variables
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())
    const { selectedLocation } = storeToRefs(selectedLocationStore())
    const { teamCards } = storeToRefs(teamCardsStore())

    //Stateful methods
    const { setCards } = scoreCardsStore()
    const { setSelectedLocation } = selectedLocationStore()
    const { addTeamCard } = teamCardsStore()

    const isLoading = ref<boolean>(true)
    const options = ref<string[]>([])
    let selectedLocationName = ref<string>("")
    let locations: Location[] = []

    const optionClicked = async (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        try {
            const locationResponse = await scoreBoardApi.get<Location>(`/get-location/${selectedValue}`)
            
            setCards(locationResponse.data.users)     
            setSelectedLocation(locationResponse.data)               
        } catch (error) {
            console.log("err in clicking option:", error);
        }
    }

    const addTeam = () => {
        addTeamCard({
            team_name: selectedLocationName.value,
            score: 0,
            players: []
        })

        // options.value = locations.filter(
        //     location => !teamCards.value.some(team => team.team_name === location.location_name)
        // ).map(location => location.location_name)
    }

    // watch(options, (newOptions) => {

    // })

    // computed(() => {

    // })

    onMounted(async () => {
        try {
            const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")
            // let locations: Location[] = []

            //Assign the response from the server to the above variable to be referenced later.
            locations = locationsResponse.data

            //create new game and add new user have different menu options.
            if (currentButtonGroupState.value === ButtonState.CREATE_NEW_TEAM_GAME) {
                // options.value = locations.filter(
                //     location => !teamCards.value.some(team => team.team_name === location.location_name)
                // ).map(location => location.location_name)
            } else {
            }
            
            //Take all of the names from the all of the locations, and assign them to the options variable to
            //listed on the dropdown menu.
            options.value = locations.map(location => {          
                return location.location_name
            }) 
            //IF: If there is no current location set, set it now, and the current name for the options dropdown.
            //ELSE: Assign the current name of the selectedLocation type to the selectedLocationName
            //string variable so it can show on the options dropdown.
            if (!selectedLocation.value) {
                selectedLocationName.value = options.value[0]
                setSelectedLocation(locations[0])
            }else{
                selectedLocationName.value = selectedLocation.value.location_name
            }

            //If there are no current cards set, and there is a selectedLocation, set the cards for the location.
            if (!scoreCards.value && selectedLocation.value) {
                setCards(selectedLocation.value.users)
            }

            console.log("Cards:", scoreCards.value);
            
        } catch (error) {
            console.log("err:", error)
        }

        isLoading.value = false
    })
</script>

<template>
    <div  :class="`location ${isDarkMode ? 'dark-mode-text' : 'light-mode-text'}`">
        <div class="current-location" v-if="currentButtonGroupState != ButtonState.CREATE_NEW_TEAM_GAME">
            Current Location: 
        </div> 
        <Icon icon="svg-spinners:180-ring" v-if="isLoading"/>
        <select 
            v-else
            v-model="selectedLocationName"
            name="locations" 
            id="locations" 
            :class="`${isDarkMode ? 'dark-mode-select' : 'light-mode-select'}`" 
            @change="optionClicked"
        >
            <option v-for="(option, index) in options" :value="option" :key="index">
                {{ option }}
            </option>
        </select>   
        <button 
            @click="addTeam"
            v-if="currentButtonGroupState === ButtonState.CREATE_NEW_TEAM_GAME" 
            class="add-team-button"
        
        >Add Team</button> 
    </div>
</template>

<style scoped lang="scss">
    .location{
        text-align: center;
        
        @media screen and (min-width: 768px) {
            display: flex;
            align-items: center;
            justify-content: center;
        }

        font-size: 30px;
        transition: 0.5s;
        
        .underline{
            text-decoration: underline;
        }

        .current-location{
            margin-right: 10px;
        }

        select {
            font-size: 18px;
            padding: 2px 5px;
            transition: 0.5s;
    
            @media screen and (min-width: 768px) {
                font-size: 28px;
            }
        }

        .add-team-button{
            padding: 10px 25px;
            margin: 0px 10px;

            border: none;
            border-radius: 5px;

            background-color: dodgerblue;
            color: aliceblue;
            font-size: 20px;
            transition: 0.3s;

            @media screen and (min-width: 768px) {
                &:hover{
                    cursor: pointer;
                    font-size: 26px;
                }
            }
        }

        .dark-mode-select{
            background-color: var(--main-bg-color);
            color: var(--primary-color);
            border: 2px solid var(--primary-color);
        }

        .light-mode-select{
            background-color: var(--primary-color);
            color: var(--main-bg-color);
            border: 2px solid var(--main-bg-color);
        }
    }
</style>