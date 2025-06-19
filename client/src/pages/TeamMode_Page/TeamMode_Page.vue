<script setup lang="ts">
    //libraries and stores
    import { teamCardsStore } from "../../stores/teamCardsStore"

    //Containers
    import TeamCardsContainer from "../../container/TeamCardsContainer/TeamCardsContainer.vue";
    import AddNewTeamContainer from "../../container/AddNewTeamContainer/AddNewTeamContainer.vue";
    
    //Components
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';

     //Stateful method(s)
    const { resetAllPoints } = teamCardsStore()
</script>

<template>
    <PageTitle titleName="Teams to Add"/>

    <!-- Contains the ADAPT locations in the select tag -->
    <AddNewTeamContainer />

    <!-- Button to reset all of the points for every team playing -->
    <div class="reset-points-wrapper">
        <ResetPoints :resetAllPoints="resetAllPoints"/>
    </div>

    <!-- Show all team cards added by the user  --> 
    <TeamCardsContainer />

    <!-- Saves a game to the server -->
    <SaveGame />
</template>

<style lang="scss">
    .reset-points-wrapper{
        margin: 20px;
    }
</style>