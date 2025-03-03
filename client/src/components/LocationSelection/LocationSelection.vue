<script setup lang="ts">
    import { darkModeStore } from "../../stores/darkModeStore"
    import { onMounted, ref } from 'vue';
    import { buttonActiveStore, ButtonState } from "../../stores/buttonActiveStore"
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
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
    const { teamCards } = storeToRefs(teamCardsStore())
    const { allLocationOptions, remainingLocationOptions } = storeToRefs(optionsStore())

    //Stateful methods
    const { setCards } = scoreCardsStore()
    const { setSelectedLocation } = selectedLocationStore()
    const { addTeamCard } = teamCardsStore()
    const { setRemainingLocationOptions, setAllLocationOptions } = optionsStore()

    const isLoading = ref<boolean>(true)
    let locations: Location[] = []
    // const options = ref<string[]>([])
    // let selectedLocationName = ref<string>("")

    const optionClicked = async (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        try {
            const locationResponse = await scoreBoardApi.get<Location>(`/get-location/${selectedValue}`)
            
            setCards(locationResponse.data.users)     
            setSelectedLocation(locationResponse.data)               
        } catch (error) {
            console.log("err in clicking option:", error);
        }

        console.log("selectedLocationName:", selectedLocationName.value);
    }

    const addTeam = () => {
        addTeamCard({
            team_name: selectedLocationName.value,
            score: 0,
            players: []
        })

        //After adding the card, remove it from options.
        setRemainingLocationOptions(locations.filter(
            location => !teamCards.value.some(team => team.team_name === location.location_name)
        ).map(location => location.location_name))
        
        selectedLocationName.value = remainingLocationOptions.value[0]
    }

    // watch(options, (newOptions) => {

    // })

    // computed(() => {

    // })

    onMounted(async () => {
        try {
            const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")

            //Assign the response from the server to the above variable to be referenced later.
            locations = locationsResponse.data

            //if the user is on "CREATE_NEW_TEAM_GAME", load the remaining locations that have not been 
            //added to the game.
            if (currentButtonGroupState.value === ButtonState.CREATE_NEW_TEAM_GAME) {
                setRemainingLocationOptions(locations.filter(
                    location => !teamCards.value.some(team => team.team_name === location.location_name)
                ).map(location => location.location_name))
            } else {
                //Take all of the names from the all of the locations, and assign them to the options 
                //variable to be listed on the dropdown menu.
                setAllLocationOptions(locations.map(location => {          
                    return location.location_name
                }))
            }            

            //IF: If there is no current location set, set it now, and the current name for the options dropdown.
            //ELSE: Assign the current name of the selectedLocation type to the selectedLocationName
            //string variable so it can show on the options dropdown.
            if (!selectedLocation.value.id) {
                selectedLocationName.value = allLocationOptions.value[0]
                setSelectedLocation(locations[0])
                console.log("selectedLocation:", selectedLocation.value);
            }else if (currentButtonGroupState.value === ButtonState.ADD_NEW_USER) {
                selectedLocationName.value = selectedLocation.value.location_name
            }else{
                selectedLocationName.value = remainingLocationOptions.value[0]   
            }

            //If there are no current cards set, and there is a selectedLocation, set the cards for the location.
            if (!scoreCards.value.length && selectedLocation.value) {
                setCards(selectedLocation.value.users)
            }

            console.log("selectedLocationName:", selectedLocationName.value);
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

        <!-- If the user is creating a new team game, choose this select tag. -->
        <Select 
            v-else-if="!isLoading && (remainingLocationOptions.length && currentButtonGroupState === ButtonState.CREATE_NEW_TEAM_GAME)"
            :options="remainingLocationOptions"
            :optionClicked="optionClicked"
        />
        
        <!-- Otherwise choose this one -->
        <Select 
            v-else-if="!isLoading && currentButtonGroupState == ButtonState.ADD_NEW_USER"
            :options="allLocationOptions"
            :optionClicked="optionClicked"
        />
        
        <!-- <select 
            v-else-if="!isLoading && (remainingLocationOptions.length && currentButtonGroupState === ButtonState.CREATE_NEW_TEAM_GAME)"
            v-model="selectedLocationName"
            name="remaining-locations" 
            id="remaining-locations" 
            :class="`${isDarkMode ? 'dark-mode-select' : 'light-mode-select'}`" 
            @change="optionClicked"
        >
            <option v-for="(option, index) in remainingLocationOptions" :value="option" :key="index+1">
                {{ option }}
            </option>
        </select>    -->

        <!-- <select 
            v-else-if="!isLoading && currentButtonGroupState == ButtonState.ADD_NEW_USER"
            v-model="selectedLocationName"
            name="locations" 
            id="locations" 
            :class="`${isDarkMode ? 'dark-mode-select' : 'light-mode-select'}`" 
            @change="optionClicked"
            >
            <option 
                v-for="(option, index) in allLocationOptions" 
                :value="option"
                :key="index"
            >
                {{ option }}
            </option>
        </select>  -->
        <button 
            @click="addTeam"
            v-if="currentButtonGroupState === ButtonState.CREATE_NEW_TEAM_GAME && remainingLocationOptions.length && !isLoading" 
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
    }
</style>