<script setup lang="ts">
    import ScoreCards from '../../components/ScoreCards/ScoreCards.vue';
    import TeamCards from '../../components/TeamCards/TeamCards.vue';
    import GameButtonGroup from '../../components/GameButtonGroup/GameButtonGroup.vue';
    import AddUserToLocation from '../../components/AddUserToLocation/AddUserToLocation.vue';
    import LocationSelection from '../../components/LocationSelection/LocationSelection.vue';
    import { onMounted, ref } from 'vue';
    import { darkModeStore } from "../../stores/darkModeStore"
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
    import { buttonActiveStore, ButtonState } from '../../stores/buttonActiveStore';
    import { storeToRefs } from 'pinia';
    import { scoreBoardApi } from "../../api/api"
    import { SavedGame } from "../../types/types"

    const isLoading = ref<boolean>(true)

    //Stateful methods
    const { resetAllPoints, getWinner, totalPoints } = scoreCardsStore()

    //Stateful variables
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())
    const { selectedLocation } = storeToRefs(selectedLocationStore())
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())

    const addSavedGame = async () => {
        try {

            //If the selectedLocation object is NOT NULL, create a new saved game.
            if (selectedLocation.value) {
                const savedGame: SavedGame = {
                    id: "",
                    winner: getWinner(),
                    location: {
                        id: "",
                        location_name: selectedLocation.value?.location_name,
                        users: scoreCards.value
                    },
                    total_points: 0,
                    average_points: 0.0,
                    created_at: new Date().toLocaleDateString()
                }
    
                const res = await scoreBoardApi.post("/save-game", savedGame)
    
                console.log("res: ", res.data)
            }

        } catch (error) {
            console.log("err:", error);
        }
    }

    onMounted(async () => {
        // if (locationsFromLocalStorage.value.length) {
        //     locations = locationsFromLocalStorage.value
        //     options.value = locations.map(location => {
        //         return location.location_name
        //     })   
        // } else {
        // }
        try {
            //Find the target location in the list of locations from the database.
            // const targetLocation = locations.find(location => {
            //     return location.location_name === selectedLocation.value
            // })

            

            // //If the scoreCards array is empty, populate it with the users from the location.
            // if (targetLocation && scoreCards.value.length === 0) {
            //     setCards(targetLocation.users)
            // }

            // //If the list of users from the target location is greater than the list of users currently,
            // //append the difference to the current list of users
            // if (targetLocation && targetLocation.users.length > scoreCards.value.length) {
            //     const index = targetLocation.users.length - scoreCards.value.length

            //     scoreCards.value = [...scoreCards.value, ...targetLocation.users.slice(index)]
            // }

            // //If the list of users from the target location is greater than the list of users currently,
            // //append the difference to the current list of users
            // if (targetLocation && targetLocation.users.length <= scoreCards.value.length) {
            //     setCards(targetLocation.users)                
            // }

            // console.log("target location:", targetLocation, "scoreCards:", scoreCards.value);
            
            //If there are no current score cards set in local storage, retrieve them from the database
            // locationsResponse.data.forEach((location, i) => {
            //     if (location.location_name === selectedLocation.value) {
            //         setCards(locationsResponse.data[i].users)
            //     }
            // })
            // if (!scoreCards) {
            // } else{
            //     console.log("card:", scoreCards.value);
                
            // }         
        } catch (error) {
            console.log("err:", error);
        }
        
        // console.log("locations:", locations, "option:", options, "host:", window.location.hostname);
        // console.log("api:", scoreBoardApi.getUri());
        
        isLoading.value = false
    })
</script>

<template>
    <div class="title">Digital Score Board</div>
    
    <!-- Contains the ADAPT locations in the select tag -->
    <LocationSelection />

    <!-- Contains the three buttons needed to select the options. -->
    <GameButtonGroup />

    <!-- Only show this when the "Create new game" or "Add new user" is clicked -->
    <AddUserToLocation />

    <div class="reset-all-points-wrapper">
        <button type="button" @click="resetAllPoints" >
            Reset All Points
        </button>
    </div>

    <div 
        v-if="currentButtonGroupState === ButtonState.ADD_NEW_USER" 
        class="total-points"
    >
        Total Points: {{ totalPoints() }} 
    </div>

    <!-- Shows Create new game form when "Create new game" button is clicked  -->
    <div 
        class="create-new-game"
        v-if="currentButtonGroupState == ButtonState.CREATE_NEW_GAME"
    >

    </div>

    <!-- Shows all users when "Add new users" is clicked -->
    <ScoreCards v-if="currentButtonGroupState === ButtonState.ADD_NEW_USER" />

    <!-- Shows Create new Team game form when "Create new team game" button is clicked  --> 
    <TeamCards v-if="currentButtonGroupState === ButtonState.CREATE_NEW_TEAM_GAME" />
 
    <div class="save-wrapper">
        <button type="button" @click="addSavedGame" :class="`${isDarkMode ? 'dark-mode' : 'light-mode'}`">
            Save Game
        </button>
    </div>
</template>

<style scoped lang="scss">
    .dark-mode{
        border: 2px var(--primary-color) solid;
        color: var(--main-text-color);
        background-color: transparent;
    }

    .light-mode{
        color: var(--primary-color);
        background-color: var(--main-text-color)
    }

    .dark-mode-text{
        color: var(--primary-color);
    }

    .light-mode-text{
        color: var(--main-bg-color);
    }

    .title{
        color: var(--main-text-color);
        text-align: center;
        font-size: 40px;
        margin: 40px 0px;

        @media screen and (min-width: 768px) {
            font-size: 60px;
            margin: 10px 0px;
        }
    }

    .error{
        text-align: center;
        color: red;
        font-size: 25px;
    }

    .reset-all-points-wrapper{
        text-align: center;
        margin: 25px;

        button{
            border-radius: 10px;
            border: none;

            padding: 10px 18px;
            font-size: 16px;
            font-weight: 700;
            transition: 0.3s;
            
            background-color: red;
            color: aliceblue;

            &:hover{
                cursor: pointer;
                background-color: rgba($color: #e30909, $alpha: .8);
            }
            
            &:active{
                transform: translateY(-5px);
            }
        }
    }

    .total-points{
        text-align: center;
        font-size: 30px;
        color: var(--main-text-color);
    }

    .save-wrapper{
        text-align: center;
        margin: 40px;

        button{
            border-radius: 10px;
            padding: 15px 35px;
    
            transition: 0.5s;
            font-size: 25px;

            &:hover{
                cursor: pointer;
                background-color: var(--main-text-color-transparent);
            }

            &:active {
                transform: translateY(-5px);
            }
        }
    }

    .save-wrapper button:active {
        transform: translateY(-5px);
    }
</style>