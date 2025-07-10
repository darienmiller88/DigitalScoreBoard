<script setup lang="ts">
    import ScoreCards from '../../components/ScoreCards/ScoreCards.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';
    import CreateNewGameContainer from '../../container/CreateNewGameContainer/CreateNewGameContainer.vue';
    import AvailablePlayersContainers from '../../container/AvailablePlayersContainers/AvailablePlayersContainers.vue';

    import { onMounted, ref } from 'vue';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"

    //Stateful methods
    const { resetAllPoints, totalPoints } = scoreCardsStore()
    let isGameCreated = ref<boolean>(false)  
  
    onMounted(async () => {
        console.log(isGameCreated);
        
        try {
           
        } catch (error) {
            console.log("err:", error);
        }
    })
</script>

<template>
    <PageTitle :titleName="'Create New Game'"/>

    <div v-if="!isGameCreated">
        <CreateNewGameContainer />
    
        <AvailablePlayersContainers />

        <!-- When this is created, add the current players to the score cards store -->
        <div class="button-wrapper">
            <button>Create Game!</button>
        </div>
    </div>
    <div v-else>
        <ResetPoints :resetAllPoints="resetAllPoints"/>
    
        <div class="total-points">
            Total Points: {{ totalPoints() }} 
        </div>
    
        <!-- Shows all users when "Add new users" is clicked -->
        <ScoreCards />
    
        <!-- Saves a game to the server -->
        <SaveGame />
    </div>
</template>

<style scoped lang="scss">
    .total-points{
        text-align: center;
        font-size: 30px;
        color: var(--primary-color);
        padding: 20px;
    }  

    .button-wrapper{
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