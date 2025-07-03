<script setup lang="ts">
    import ScoreCards from '../../components/ScoreCards/ScoreCards.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';
    import CreateNewGameContainer from '../../container/CreateNewGameContainer/CreateNewGameContainer.vue';
    import AvailablePlayersContainers from '../../container/AvailablePlayersContainers/AvailablePlayersContainers.vue';

    import { onMounted } from 'vue';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"

    //Stateful methods
    const { resetAllPoints, totalPoints } = scoreCardsStore()
  
    onMounted(async () => {
        try {
           
        } catch (error) {
            console.log("err:", error);
        }
    })
</script>

<template>
    <PageTitle :titleName="'Create New Game'"/>

    <CreateNewGameContainer />

    <AvailablePlayersContainers />

    <ResetPoints :resetAllPoints="resetAllPoints"/>

    <div class="total-points">
        Total Points: {{ totalPoints() }} 
    </div>

    <!-- Shows all users when "Add new users" is clicked -->
    <!-- <ScoreCards /> -->

    <!-- Saves a game to the server -->
    <!-- <SaveGame /> -->
</template>

<style scoped lang="scss">
    .error{
        text-align: center;
        color: red;
        font-size: 25px;
    }

    .total-points{
        text-align: center;
        font-size: 30px;
        color: var(--primary-color);
        padding: 20px;
    }  
</style>