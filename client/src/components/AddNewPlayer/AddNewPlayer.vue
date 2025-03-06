<script setup lang="ts">
import { onMounted } from 'vue';

    
    const props = defineProps<{
        players: string[]
        locationName: string
    }>()

    onMounted(() => {
        console.log("AddNewPlayer props:", props.players, "and", props.locationName);
    })
</script>

<template>
    <!-- <ul v-for="(player, index) in players">
         <li :key="index">{{ player }}</li>
    </ul> -->
    <div>Players in {{ props.locationName }}</div>
    <select name="team-players" id="team-players">
        <option v-for="(player, index) in props.players" :value="player" :key="index">
            {{ player }}
        </option>
    </select>
    <!-- <div class="add-new-player">
        <h1>{{ props.playerName }}</h1>
    </div> -->
</template>

<style scoped>

</style>