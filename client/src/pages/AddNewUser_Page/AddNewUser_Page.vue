<script setup  lang="ts">
    import AddUserToLocationContainer from '../../container/AddUserToLocationContainer/AddUserToLocationContainer.vue';
    import UserCardsContainer from '../../container/UserCardsContainer/UserCardsContainer.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';

    import { setLocation, currentLocation } from "./currentLocation"
</script>

<template>
    <PageTitle :titleName="'Add New Player'"/>

    <AddUserToLocationContainer :changeLocation="setLocation" :currentLocation="currentLocation"/>

    <UserCardsContainer :currentLocation="currentLocation"/>
</template>
