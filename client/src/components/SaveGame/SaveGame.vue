<script setup lang="ts">
    import { ref } from 'vue';
    import { storeToRefs } from 'pinia';
    import { SavedGame } from '../../types/types';
    import { teamCardsStore } from '../../stores/teamCardsStore';
    import { scoreBoardApi } from '../../api/api';
    import { buttonActiveStore, ButtonState } from '../../stores/buttonActiveStore';
    import Loading from '../Loading/Loading.vue';
    import Modal from '../Modal/Modal.vue';

    //Stateful methods
    const { getWinningTeam, getTotalPoints, getAveragePoints, getPlayers } = teamCardsStore()

    //Stateful variables
    const { teamCards } = storeToRefs(teamCardsStore())
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())

    let isLoading = ref<boolean>(false)

    const addSavedGame = async () => {
        isLoading.value = false

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

        isLoading.value = true
    }
</script>

<template>
   <div class="save-wrapper">
        <button type="button" @click="addSavedGame" disabled>
            <Loading :height="30" :usePrimary="false" v-if="isLoading"/>
            <div v-else> Save Game </div>
        </button>
    </div>

    <!-- <Modal 
    
    /> -->
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