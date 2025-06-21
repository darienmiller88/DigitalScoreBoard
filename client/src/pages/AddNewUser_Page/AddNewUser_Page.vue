<script setup  lang="ts">
    import AddUserToLocationContainer from '../../container/AddUserToLocationContainer/AddUserToLocationContainer.vue';
    import UserCardsContainer from '../../container/UserCardsContainer/UserCardsContainer.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
    import { ref } from 'vue';

    let currentLocationRef = ref<string>("")
</script>

<template>
    <PageTitle :titleName="'Add New Player'"/>

    <AddUserToLocationContainer :currentLocation="currentLocationRef"/>

    <UserCardsContainer :currentLocation="currentLocationRef"/>
</template>
