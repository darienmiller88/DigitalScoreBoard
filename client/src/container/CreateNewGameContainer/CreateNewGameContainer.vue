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
    <div>
        CreateNewGameContainer
        <Select :options="allLocationOptions" :selectModel="currentLocation" :onChange="onChangeSelect" />
    </div>
</template>

<style scoped lang="scss">

</style>