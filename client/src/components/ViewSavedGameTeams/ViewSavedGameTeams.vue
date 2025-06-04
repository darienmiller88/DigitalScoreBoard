<script setup lang="ts">
    import { Team } from "../../types/types"

    const props = defineProps<{
        teamsInSavedGame: Team[]
    }>()
</script>

<template>
    <div class="teams"> 
        <div class="team" v-for="team in props.teamsInSavedGame">
            {{ team.team_name }}

            
        </div>
    </div>
</template>

<style scoped lang="scss">
    .teams{
        border: 2px solid black;

        .team{
            border: 2px solid red;
        }
    }
</style>