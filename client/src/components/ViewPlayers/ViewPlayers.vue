<script setup lang="ts">
    const props = defineProps<{
        teamPlayers: string[]
    }>()
</script>

<template>
    <ul v-for="(player, index) in props.teamPlayers">
         <li :key="index">{{ player }}</li>
    </ul>
</template>

<style scoped>

</style>