<script setup lang="ts">
    const props = defineProps<{
        titleName: string
    }>()
</script>

<template>
    <div class="title">{{ props.titleName }}</div>
</template>

<style scoped lang="scss">
    .title{
        color: var(--main-text-color);
        text-align: center;
        font-size: 40px;
        margin: 40px 0px;

        @media screen and (min-width: 768px) {
            font-size: 60px;
            margin: 10px 0px;
        }
    }
</style>