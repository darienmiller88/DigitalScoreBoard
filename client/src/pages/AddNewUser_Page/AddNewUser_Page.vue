<script setup lang="ts">
    import AddUserToLocationContainer from '../../container/AddUserToLocationContainer/AddUserToLocationContainer.vue';
    import UserCardsContainer from '../../container/UserCardsContainer/UserCardsContainer.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
</script>

<template>
    <PageTitle :titleName="'Add New Player'"/>

    <AddUserToLocationContainer />

    <UserCardsContainer />
</template>
