<script setup lang="ts">
    import { ref } from "vue";
    import { storeToRefs } from "pinia";
    import { darkModeStore } from "../../stores/darkModeStore"
    import { buttonActiveStore, ButtonState } from "../../stores/buttonActiveStore"
    import { optionsStore } from "../../stores/optionsStore";
    import { useWindowSize } from "@vueuse/core";
    import { teamCardsStore } from "../../stores/teamCardsStore";
    import { scoreBoardApi } from "../../api/api";
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
    import { Location } from "../../types/types";

    const { width } = useWindowSize();
    let islocationsLoaded = ref<boolean>(false)
    let locations: Location[] = []

    //Stateful variables
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { teamCards } = storeToRefs(teamCardsStore())
    const{ currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const { selectedLocationName, selectedLocation } = storeToRefs(selectedLocationStore())
    const { remainingLocationOptions } = storeToRefs(optionsStore())
    
    //Stateful methods
    const { setButtonActive } = buttonActiveStore()
    const { setAllLocationOptions, setRemainingLocationOptions } = optionsStore()

    const loadLocationOptions = async (buttonState: ButtonState) => {
        //If the locations are NOT loaded yet, load them, and set the loaded flag to true to prevent
        //unnecessary server calls.
        if(!islocationsLoaded.value){
            try {
                const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")
    
                locations = locationsResponse.data
                islocationsLoaded.value = true
            } catch (error) {
                console.log("err:", error)
            }
        }

        selectedLocationName.value = selectedLocation.value.location_name
        if(buttonState === ButtonState.CREATE_NEW_TEAM_GAME) {
            //When "CREATE_NEW_TEAM_GAME" is clicked, set options to the locations that have NOT
            //been added as a team card.
            setRemainingLocationOptions(locations.filter(
                location => !teamCards.value.some(team => team.team_name === location.location_name)
            ).map(location => location.location_name))

            selectedLocationName.value = remainingLocationOptions.value[0]            
            setButtonActive(ButtonState.CREATE_NEW_TEAM_GAME)
        }else if(buttonState == ButtonState.ADD_NEW_USER){
            //load all locations from the server, and add them to options.
            setAllLocationOptions(locations.map(location => {          
                return location.location_name
            }))

            setButtonActive(ButtonState.ADD_NEW_USER)
        }else{
            setButtonActive(ButtonState.CREATE_NEW_GAME)
        }
    }
</script>

<template>
    <div class="button-group">
        <button 
            :class="`
                ${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}
                ${currentButtonGroupState == ButtonState.CREATE_NEW_GAME ? 'active' : ''}
            `" 
            @click="() => loadLocationOptions(ButtonState.CREATE_NEW_GAME)"
        >Create new game</button>
        <span v-if="width >= 768" :class="`${isDarkMode ? 'dark-mode-span' : 'light-mode-span'}`"></span>
        <button 
            :class="`
                ${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}
                ${currentButtonGroupState == ButtonState.ADD_NEW_USER ? 'active' : ''}
            `"
            @click="() => loadLocationOptions(ButtonState.ADD_NEW_USER)"
        >Add new user</button>
        <span v-if="width >= 768" :class="`${isDarkMode ? 'dark-mode-span' : 'light-mode-span'}`"></span>
        <button 
            :class="`
                ${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}
                ${currentButtonGroupState == ButtonState.CREATE_NEW_TEAM_GAME ? 'active' : ''}
            `"
            @click="() => loadLocationOptions(ButtonState.CREATE_NEW_TEAM_GAME)"
        >Create new team game</button>
    </div>
</template>

<style scoped lang="scss"> 
    .button-group{
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: stretch;
        
        padding: 0px 15px;
        margin-top: 30px;

        @media screen and (min-width: 768px) {
            flex-direction: row;
            padding: 0px;
        }

        span{
            margin: 0px 15px;
            transition: 0.3s;
        }

        .dark-mode-span{
            border: 1px solid aliceblue;
        }

        .light-mode-span{
            border: 1px solid black;
        }

        .active{
            background-color: dodgerblue;
        }

        button{
            background-color: black;
            color: aliceblue;
            border: none;
            border-radius: 8px;

            padding: 15px 28px;
            margin-bottom: 18px;

            font-size: 20px;
            transition: 0.3s;

            box-shadow: 0px 4px 8px rgba(100, 100, 100, 0.25), /* Bottom shadow */
                0px -4px 8px rgba(100, 100, 100, 0.25), /* Top shadow */
                4px 0px 8px rgba(100, 100, 100, 0.25), /* Right shadow */
                -4px 0px 8px rgba(100, 100, 100, 0.25);

            @media screen and (min-width: 768px) {
                margin-bottom: 0px;

                &:hover{
                    cursor: pointer;
                    padding: 15px 38px;
                    font-size: 28px;
                }
            }

            &:active{
                transform: translateY(-5px);
            }
        }
    }
</style>