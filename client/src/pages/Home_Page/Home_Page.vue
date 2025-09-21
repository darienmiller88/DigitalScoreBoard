<script setup lang="ts">
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
    import SaveGameButton from '../../components/SaveGameButton/SaveGameButton.vue';
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';
    import ScoreCards from '../../container/HomePageContainers/ScoreCardsContainer/ScoreCardsContainer.vue';
    import SelectLocationContainer from '../../container/HomePageContainers/SelectLocationContainer/SelectLocationContainer.vue';
    import AvailablePlayersContainers from '../../container/HomePageContainers/AvailablePlayersContainers/AvailablePlayersContainers.vue';

    import { ref } from 'vue';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { HomePageStore } from '../../stores/PageStores/HomePageStore';
    import { storeToRefs } from 'pinia';
    import { useToast } from "vue-toastification";
    import { SavedGame } from '../../types/types';
    import { scoreBoardApi } from '../../api/api';
    import {useWindowSize } from "@vueuse/core"

    const { width } = useWindowSize()

    //ref variable
    const { currentPlayersInGame, currentLocation, isGameCreated, remainingPlayersInGame } = storeToRefs(HomePageStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())

    //Stateful methods
    const { resetAllPoints, totalPoints, sortPlayersByScore } = scoreCardsStore()
    const { toggleGameCreatedStatus } = HomePageStore()
    const toast = useToast()

    let isLoading = ref<boolean>(false)
    let isMinimumNumberOfPlayerErrorShown = ref<boolean>(false)

    const createGame = () => {
        
        //If the client added no players, prevent the game from being created and flash an error.
        if (currentPlayersInGame.value.length == 0) {
            isMinimumNumberOfPlayerErrorShown.value = true

            setTimeout(() => {
                isMinimumNumberOfPlayerErrorShown.value = false
            }, 2000);
        } else {
            toggleGameCreatedStatus(true)

            scoreCards.value = currentPlayersInGame.value.map(playerName => ({
                username: playerName,
                score: 0
            }))
        }
    }

    const addMorePlayers = () => {
        toggleGameCreatedStatus(false)
    }

    const returnToGame = () => {
        toggleGameCreatedStatus(true)
    }

    const closeGame = () => {
        scoreCards.value = []
        currentPlayersInGame.value = []
        isGameCreated.value = false
    }

    const endAndSaveGame = async () => {
        let savedGame: SavedGame = {
            id: '',
            winner: {
                username: '',
                score: 0
            },
            location_name: currentLocation.value,
            created_at: new Date().toISOString(),
            total_points: 0,
            average_points: 0, 
            players: scoreCards.value
        }
        
        isLoading.value = true
        try {
            await scoreBoardApi.post<SavedGame>("/save-game", savedGame)
            
            toast.success("Game saved!", { timeout: 2500 })
            scoreCards.value = []
        } catch (error) {
            console.log(error);
        }
        
        isLoading.value = false
        closeGame()
    }
</script>

<template>
    
    <!-- If there are no cards, or the flag for when a game has been made has not been set to true, render this block  -->
    <div v-if="!scoreCards.length || !isGameCreated">
        <PageTitle :titleName="'Create New Game'"/>

        <SelectLocationContainer v-if="!scoreCards.length"/>
    
        <AvailablePlayersContainers />

        <!-- if the client did not add any players, flash this error -->
        <div class="error" v-if="isMinimumNumberOfPlayerErrorShown">Please add at least one player to the game.</div>
       
        <!-- Otherwise, show the create game button -->
        <div class="create-game-button-wrapper" v-else>
            <button @click="createGame" v-if="!scoreCards.length">
                Create Game!
            </button>

            <!-- SHow this when the user is trying to add more players to a game already started -->
            <button @click="returnToGame" v-else>
                Return to Game
            </button>
        </div>
    </div>

    <!-- This section will be rendered if the client has already created a game yet. -->
    <div v-else>
        <PageTitle :titleName="'Game in session'"/>

        <div class="total-points">
            Total Points: {{ totalPoints() }} 
        </div>

        <div class="reset-sort-wrapper">
            <ResetPoints :resetAllPoints="resetAllPoints"/>
    
            <div class="sort-players-wrapper">
                <button @click="sortPlayersByScore">Sort Players by score</button>
            </div>
        </div>

        <!-- Shows all users when "Add new users" is clicked -->
        <ScoreCards />
        
        <div class="buttons-wrapper">
            <button class="close-game base-btn" @click="closeGame">Close game</button>

            <div class="line" v-if="width >= 768"></div>

            <!-- Saves a game to the server -->
            <div class="save-game-wrapper">
                <SaveGameButton :endAndSaveGame="endAndSaveGame" :isLoading="isLoading"/>
            </div>

            <!-- Only show this line and the below button if there are no more players to be added to the game. -->
            <div class="line" v-if="width >= 768 && remainingPlayersInGame.length"></div>
            <button class="add-more base-btn" @click="addMorePlayers" v-if="remainingPlayersInGame.length">Add more players</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .total-points{
        text-align: center;
        font-size: 30px;
        color: var(--primary-color);
        padding: 20px;
    }  

    .error{
        color: red;
        text-align: center;
        font-size: 25px;
    }

    .reset-sort-wrapper{

        @media (min-width: 768px) {
            display: flex;
            width: fit-content;
            margin: auto;
            margin-top: 20px;
            margin-bottom: 20px;
        }
    }

    .sort-players-wrapper{
        text-align: center;
        margin-bottom: 20px;
        
        button{
            padding: 10px 20px;
            background-color: var(--primary-color);
            color: var(--bg-color);
            border: none;
            border-radius: 10px;
            font-size: 18px;
            transition: 0.3s;

            @media (min-width: 3840px) {
                font-size: 60px;
                padding: 20px 36px;
            }

            &:hover{
                cursor: pointer;
                background-color: var(--primary-color-transparent);
            }
        }
    }

    .buttons-wrapper{
        display: flex;
        width: fit-content;
        flex-direction: column;
        align-items: center;

        margin: auto;
        margin-top: 20px;
        margin-bottom: 20px;

        @media (min-width: 768px) {
            flex-direction: row;
        }

        @media (min-width: 3840px) {
            margin-top: 60px;
        }

        .base-btn{
            border-radius: 12px;
            border: none;
            color: white;
    
            font-size: 24px;
            font-weight: 600;
    
            padding: 15px 20px;
            width: 65vw;
            margin: 15px 10px;
            transition: 0.3s;

            @media (min-width: 768px) {
                width: unset;
            }

            @media (min-width: 3840px) {
                width: unset;
                font-size: 60px;
                padding: 30px 40px;
            }
        }

        .close-game{
            background-color: red;

            &:hover{
                background-color: rgba(255, 0, 0, 0.5);
                cursor: pointer;
            }
        }

        .line{
            border: solid 2px var(--toggle-background);
            height: 5vh;
            margin: 5px 0px;

            @media (min-width: 1025px) {
                height: 10vh;
            }

            @media (min-width: 3840px) {
                border: solid 4px var(--toggle-background);
                height: 5vh
            }
        }

        .add-more{
            background-color: green;

            &:hover{
                background-color: rgba(0, 100, 0, 0.5);
                cursor: pointer;
            }
        }

        .save-game-wrapper{
            margin: 0px 10px;
        }
    }

    .create-game-button-wrapper{
        text-align: center;

        button{
            padding: 10px 20px;
            margin: 15px;
            border-radius: 8px;
            border: none;
            background-color: var(--primary-color);
            color: var(--bg-color);
            font-size: 22px;
            font-weight: 600;
            transition: 0.3s;
        

            &:hover{
                cursor: pointer;
                background-color: var(--primary-color-transparent);
            }
        }
    }
</style>