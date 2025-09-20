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
        color: var(--primary-color);
        text-align: center;
        font-size: 40px;
        margin: 30px 0px;

        @media screen and (min-width: 768px) {
            font-size: 60px;
            margin: 10px 0px;
        }

        @media screen and (min-width: 3840px) {
            font-size: 120px;
            margin: 20px 0px;
        }
    }
</style>