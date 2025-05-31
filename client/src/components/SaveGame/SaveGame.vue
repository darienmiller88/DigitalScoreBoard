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
        <button type="button" @click="addSavedGame">
            Save Game
        </button>
    </div>
</template>

<style scoped lang="scss">
     .save-wrapper{
        text-align: center;
        margin: 40px;

        button{
            border-radius: 10px;
            border: none;
            padding: 15px 35px;
            background-color: var(--primary-color);
            color: aliceblue;
    
            transition: 0.5s;
            font-size: 25px;

            &:hover{
                cursor: pointer;
                background-color: var(--primary-color-transparent);
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