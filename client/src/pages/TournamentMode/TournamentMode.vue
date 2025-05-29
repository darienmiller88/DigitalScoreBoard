<script setup lang="ts">
    import TeamCards from '../../components/TeamCards/TeamCards.vue';
    import AddNewTeam from '../../components/AddNewTeam/AddNewTeam.vue';
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import { teamCardsStore } from "../../stores/teamCardsStore"
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';

     //Stateful methods
    const { resetAllPoints } = teamCardsStore()
</script>

<template>
    <!-- Contains the ADAPT locations in the select tag -->
    <AddNewTeam />

    <!-- Button to reset all of the points for every team playing -->
    <ResetPoints :resetAllPoints="resetAllPoints"/>

    <!-- Shows Create new Team game form when "Create new team game" button is clicked  --> 
    <TeamCards />

    <!-- Saves a game to the server -->
    <SaveGame />
</template>

<style scoped lang="scss">
    
</style>