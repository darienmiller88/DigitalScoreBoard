<script setup lang="ts">
    import { Card } from "../../types/types"

    const props = defineProps<{
        playersInSavedGame: Card[]
    }>()
</script>

<template>
    <div>
        view saved game players
        <!-- If the saved game is a Team, render out the teams differentl -->
        <ul v-for="player in props.playersInSavedGame"> {{ player.username }}</ul>        
    </div>
</template>

<style scoped lang="scss">

</style>