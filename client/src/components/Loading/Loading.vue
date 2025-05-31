<script setup lang="ts">
    import { Icon } from "@iconify/vue"

    const props = defineProps<{
        height: number
    }>()
</script>

<template>
    <div class="icon-wrapper">
        <Icon icon="svg-spinners:180-ring"  :height="props.height" :width="props.height"/>
    </div>
</template>

<style scoped lang="scss">
    .icon-wrapper{
        width: fit-content;
        margin: auto;

        color: var(--primary-color);
    }
</style>