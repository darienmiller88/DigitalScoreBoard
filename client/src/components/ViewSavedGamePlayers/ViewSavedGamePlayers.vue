<script setup lang="ts">
    import { Card, Team } from "../../types/types"

    const props = defineProps<{
        isSavedGameATeamGame: boolean
        players: Card | Team
    }>()
</script>

<template>
    <div>
        view saved game players
        <!-- If the saved game is a Team, render out the teams differentl -->
        <ul v-if="isSavedGameATeamGame" v-for="player in props.players"> {{ player }}</ul>
        
        <!-- Otherwise, render out the list of players -->
        <ol v-else v-for="player in props.players"> {{ player }}</ol>
    </div>
</template>

<style scoped lang="scss">

</style>