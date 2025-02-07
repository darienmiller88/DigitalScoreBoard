<script setup lang="ts">
    import { darkModeStore } from "../../stores/darkModeStore"
    import { onMounted, ref,  } from 'vue';
    import { buttonActiveStore, ButtonState} from "../../stores/buttonActiveStore"
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
    import { storeToRefs } from "pinia";
    import { scoreBoardApi } from "../../api/api"
    import { Location } from "../../types/types"
    import { Icon } from "@iconify/vue"

    //Stateful variables
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())
    const { selectedLocation } = storeToRefs(selectedLocationStore())

    //Stateful methods
    const { setCards } = scoreCardsStore()
    const { setSelectedLocation } = selectedLocationStore()

    const isLoading = ref<boolean>(true)
    const options = ref<string[]>([])
    const selectedLocationName = ref<string>((selectedLocation.value) ? selectedLocation.value.location_name : '')
    let locations: Location[] = []

    const optionClicked = async (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        try {
            const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")
            
            // locationsResponse.data
            locationsResponse.data.forEach(location => {
                if(location.location_name === selectedValue){
                    setCards(location.users)     
                    setSelectedLocation(location)               
                }
            })

            console.log("options:", options.value);
            
        } catch (error) {
            console.log("err in clicking option:", error);
        }
    }

    onMounted(async () => {
        try {
            const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")
            
            locations = locationsResponse.data
            options.value = locations.map(location => {          
                return location.location_name
            })     

            //If there is no current location set, set it now.
            if (!selectedLocation.value) {
                //options.value[0]
                setSelectedLocation(locations[0])
            }

            //If there are no current cards set, and there is a selectedLocation, set the card for the location.
            if (!scoreCards.value && selectedLocation.value) {
                setCards(selectedLocation.value.users)
            }
        } catch (error) {
            console.log("err:", error)
        }

        isLoading.value = false
    })
</script>

<template>
    <div :class="`location ${isDarkMode ? 'dark-mode-text' : 'light-mode-text'}`">
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