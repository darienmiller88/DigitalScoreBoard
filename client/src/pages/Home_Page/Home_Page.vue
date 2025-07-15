<script setup lang="ts">
    import ScoreCards from '../../components/ScoreCards/ScoreCards.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';
    import CreateNewGameContainer from '../../container/CreateNewGameContainer/CreateNewGameContainer.vue';
    import AvailablePlayersContainers from '../../container/AvailablePlayersContainers/AvailablePlayersContainers.vue';

    import { onMounted, ref } from 'vue';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { HomePageStore } from '../../stores/HomePageStore';
    import { storeToRefs } from 'pinia';
    import { useToast } from "vue-toastification";

    //ref variable
    const { currentPlayersInGame } = storeToRefs(HomePageStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())

    //Stateful methods
    const { resetAllPoints, totalPoints } = scoreCardsStore()
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
            //Take the players the client added, and turn them into scorecards
            scoreCards.value = currentPlayersInGame.value.map(playerName => ({
                username: playerName,
                score: 0
            }))
    
            //set the game status to true to show the scorecards
            toggleGameCreatedStatus(true)
        }
    }

    const endAndSaveGame = () => {
        isLoading.value = true
        toast.success("Game saved!", { timeout: 2500 })

        setTimeout(() => {
            isLoading.value = false
            scoreCards.value = []
        }, 3000);
    }

    onMounted(async () => {        
        try {
           
        } catch (error) {
            console.log("err:", error);
        }
    })
</script>

<template>
    
    <div v-if="!scoreCards.length">
        <PageTitle :titleName="'Create New Game'"/>

        <CreateNewGameContainer />
    
        <AvailablePlayersContainers />

        <!-- if the client did not add any players, flash this error -->
        <div class="error" v-if="isMinimumNumberOfPlayerErrorShown">Please add at least one player to the game.</div>
       
        <!-- Otherwise, show the create game button -->
        <div class="create-game-button-wrapper" v-else>
            <button @click="createGame">Create Game!</button>
        </div>
    </div>
    <div v-else>
        <PageTitle :titleName="'Game in session'"/>

        <ResetPoints :resetAllPoints="resetAllPoints"/>
    
        <div class="total-points">
            Total Points: {{ totalPoints() }} 
        </div>
    
        <!-- Shows all users when "Add new users" is clicked -->
        <ScoreCards />
        
        <div class="buttons-wrapper">

            <button class="close-game">Close game</button>

            <div class="line"></div>

            <!-- Saves a game to the server -->
            <div class="save-game-wrapper">
                <SaveGame :endAndSaveGame="endAndSaveGame" :isLoading="isLoading"/>
            </div>
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

    .buttons-wrapper{
        display: flex;
        width: fit-content;
        margin: auto;
        margin-top: 20px;
        margin-bottom: 20px;
        // border: 2px solid red;

        .close-game{
            background-color: red;
            border-radius: 8px;
            border: none;
            color: white;
            font-size: 24px;
            padding: 10px 20px;
            margin: 0px 10px;
        }

        .line{
            border: solid 2px white;
        }

        .save-game-wrapper{
            margin: 0px 10px;
        }
    }

    .create-game-button-wrapper{
        text-align: center;

        button{
            padding: 10px 20px;
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