<script setup lang="ts">
    //libraries and stores
    import { teamCardsStore } from "../../stores/teamCardsStore"

    //Components
    // import TeamCards from '../../components/TeamCards/TeamCards.vue';
    import TeamCardsContainer from "../../container/TeamCardsContainer/TeamCardsContainer.vue";
    import AddNewTeam from '../../components/AddNewTeam/AddNewTeam.vue';
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';

     //Stateful method(s)
    const { resetAllPoints } = teamCardsStore()
</script>

<template>
    <PageTitle titleName="Teams to Add"/>

    <!-- Contains the ADAPT locations in the select tag -->
    <AddNewTeam />

    <!-- Button to reset all of the points for every team playing -->
    <ResetPoints :resetAllPoints="resetAllPoints"/>

    <!-- Show all team cards added by the user  --> 
    <TeamCardsContainer />

    <!-- Saves a game to the server -->
    <SaveGame />
</template>