<script setup lang="ts">
    import { SavedGame } from '../../types/types';

    const props = defineProps<SavedGame>()
</script>

<template>
    <div class="game">

    </div>
</template>

<style scoped lang="scss">
    .game{
        border: 3px dotted green;
    }
</style>