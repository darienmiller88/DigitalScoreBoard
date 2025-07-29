<script setup lang="ts">
    import { Icon } from "@iconify/vue"

    defineProps<{
        height: number
        usePrimary: boolean
    }>()
</script>

<template>
    <div :class="`icon-wrapper ${usePrimary ? 'primary' : 'bg'}`">
        <Icon icon="svg-spinners:180-ring" :height="height" :width="height"/>
    </div>
</template>

<style scoped lang="scss">
    .icon-wrapper{
        width: fit-content;
        margin: auto;
    }

    .primary{
        color: var(--primary-color);
    }

    .bg{
        color: var(--bg-color);
    }
</style>