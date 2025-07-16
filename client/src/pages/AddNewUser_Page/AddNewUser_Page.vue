<script setup  lang="ts">
    import AddUserToLocationContainer from '../../container/AddPlayerPageContainers/AddUserToLocationContainer/AddUserToLocationContainer.vue';
    import UserCardsContainer from '../../container/AddPlayerPageContainers/UserCardsContainer/UserCardsContainer.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';

    import { 
        setLocation, 
        currentLocation, 
        setPlayers, 
        players, 
        removePlayerFromArray, 
        addNewPlayerToArray,
        editPlayerName
    } from "./composables"
</script>

<template>
    <PageTitle :titleName="'Add New Player'"/>

    <AddUserToLocationContainer 
        :changeLocation="setLocation" 
        :currentLocation="currentLocation"
        :players="players"
        :addNewPlayerToArray="addNewPlayerToArray"    
    />
    
    <UserCardsContainer
        :players="players" 
        :setPlayers="setPlayers" 
        :removePlayerFromArray="removePlayerFromArray" 
        :currentLocation="currentLocation"
        :editPlayerName="editPlayerName"
    />
</template>
