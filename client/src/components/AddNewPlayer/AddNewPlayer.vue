<script setup lang="ts">
    const props = defineProps<{
        playerName: string
        players: string[]
    }>()
</script>

<template>
    <ul v-for="(player, index) in players">
         <li :key="index">{{ player }}</li>
    </ul>
    <!-- <div class="add-new-player">
        <h1>{{ props.playerName }}</h1>
    </div> -->
</template>

<style scoped>

</style>