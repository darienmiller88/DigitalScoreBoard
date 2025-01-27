<script setup lang="ts">
    const props = defineProps<{
        playerName: string
    }>()
</script>

<template>
    <div class="add-new-player">
        <h1>{{ props.playerName }}</h1>
    </div>
</template>

<style scoped>

</style>