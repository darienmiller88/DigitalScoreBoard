<script setup lang="ts">
    import { storeToRefs } from 'pinia';
    import { SavedGame } from '../../types/types';
    import { teamCardsStore } from '../../stores/teamCardsStore';
    import { scoreBoardApi } from '../../api/api';
    import { darkModeStore } from "../../stores/darkModeStore"
    import { buttonActiveStore, ButtonState } from '../../stores/buttonActiveStore';

    //Stateful methods
    const { getWinningTeam, getTotalPoints, getAveragePoints, getPlayers } = teamCardsStore()

    //Stateful variables
    const { teamCards } = storeToRefs(teamCardsStore())
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())

    const addSavedGame = async () => {
        try {

            //If the selectedLocation object is NOT NULL, create a new saved game.
            if (currentButtonGroupState.value === ButtonState.CREATE_NEW_TEAM_GAME) {
                const savedGame: SavedGame = {
                    id: "",
                    winner: getWinningTeam(),
                    teams: teamCards.value,
                    location: {
                        location_name: "Lawrence",
                        users: getPlayers(),
                        id: ""
                    },
                    total_points: getTotalPoints(),
                    average_points: getAveragePoints(),
                    created_at: new Date().toLocaleDateString()
                }
                
                console.log("game:", savedGame);
                
                const res = await scoreBoardApi.post("/save-game", savedGame)
                console.log("res: ", res.data)
            }

        } catch (error) {
            console.log("err:", error);
        }
    }
</script>

<template>
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