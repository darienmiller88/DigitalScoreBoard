<script setup lang="ts">
    import { Team } from "../../types/types"

    const props = defineProps<{
        teamsInSavedGame: Team[]
    }>()
</script>

<template>
    <div>
        view saved game teams

        <ul v-for="team in props.teamsInSavedGame"> {{ team.team_name }}</ul>        
    </div>
</template>

<style scoped lang="scss">

</style>