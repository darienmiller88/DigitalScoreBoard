<script setup lang="ts">
    import Select from '../../components/Select/Select.vue';
    
    import { optionsStore } from '../../stores/optionsStore';
    import { HomePageStore } from '../../stores/HomePageStore';
    import { storeToRefs } from 'pinia';
    import { onMounted } from 'vue';

    const { currentLocation } = storeToRefs(HomePageStore())
    const { setCurrentLocation } = HomePageStore()
    const { allLocationOptions } = optionsStore()


    const onChangeSelect = async (event: Event) => {
        setCurrentLocation((event.target as HTMLSelectElement).value);
    }
    
    onMounted(() => {
        //Load the first location in if there is no current location saved in local storage
        if (!currentLocation.value.length) {
            currentLocation.value = allLocationOptions[0]
        }
    })
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