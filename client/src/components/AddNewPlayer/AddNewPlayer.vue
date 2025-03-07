<script setup lang="ts">
import { onMounted } from 'vue';

    
    const props = defineProps<{
        teamPlayers: string[]
        locationName: string
    }>()

    onMounted(() => {
        console.log("AddNewPlayer props:", props.teamPlayers, "and", props.locationName);
    })
</script>

<template>
    <h2 class="players-location">Players in {{ props.locationName }}:</h2>
    <select name="team-players" id="team-players">
        <option v-for="(player, index) in props.teamPlayers" :value="player" :key="index">
            {{ player }}
        </option>
    </select>
</template>

<style scoped lang="scss">
    .players-location{
        text-align: center;
    }

    select{
        margin: auto;
    }
</style>