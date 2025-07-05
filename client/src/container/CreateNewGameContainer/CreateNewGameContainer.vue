<script setup lang="ts">
    import Select from '../../components/Select/Select.vue';
    
    import { ref } from 'vue';
    import { optionsStore } from '../../stores/optionsStore';

    const { allLocationOptions } = optionsStore()

    let currentLocation = ref<string>(allLocationOptions[0])

    const onChangeSelect = (event: Event) => {
        currentLocation.value = (event.target as HTMLSelectElement).value;
    }
</script>

<template>
    <div class="choose-site-wrapper">
        <span>Choose site:</span>
        <Select :options="allLocationOptions" :selectModel="currentLocation" :onChange="onChangeSelect" />
    </div>
</template>

<style scoped lang="scss">
    .choose-site-wrapper{
        text-align: center;
        margin: 20px;

        span{
            font-size: 24px;
            margin-right: 10px;
            color: var(--primary-color);

            @media (min-width: 768px) {
                font-size: 30px;
            }
        }
    }
</style>