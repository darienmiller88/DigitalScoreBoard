<script setup lang="ts">
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
    import SavedGames from '../../components/SavedGames/SavedGames.vue';
</script>

<template>
    <!-- Display title -->
    <PageTitle :titleName="'Saved Games'"/>
    
    <!-- Display all of the games the User has stored -->
    <SavedGames />
</template>