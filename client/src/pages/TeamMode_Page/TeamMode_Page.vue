<script setup lang="ts">
    import TeamCards from '../../components/TeamCards/TeamCards.vue';
    import AddNewTeam from '../../components/AddNewTeam/AddNewTeam.vue';
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';
    import { teamCardsStore } from "../../stores/teamCardsStore"

     //Stateful method(s)
    const { resetAllPoints } = teamCardsStore()
</script>

<template>
    <!-- Contains the ADAPT locations in the select tag -->
    <AddNewTeam />

    <!-- Button to reset all of the points for every team playing -->
    <ResetPoints :resetAllPoints="resetAllPoints"/>

    <!-- Show all team cards added by the user  --> 
    <TeamCards />

    <!-- Saves a game to the server -->
    <SaveGame />
</template>