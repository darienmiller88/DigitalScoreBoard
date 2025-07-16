<script setup lang="ts">
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
    import SavedGamesContainer from '../../container/ViewGamesPageContainers/SavedGamesContainer/SavedGamesContainer.vue';
</script>

<template>
    <!-- Display title -->
    <PageTitle :titleName="'Saved Games'"/>
    
    <!-- Display all of the games the User has stored -->
    <SavedGamesContainer />
</template>